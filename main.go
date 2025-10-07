package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/logger"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/tracer"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/pago"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/pago/repository/pagodb"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/reserva"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/reserva/repository/reservadb"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario/repository/usuariodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var build = "dev"

func main() {
	var (
		ctx = context.Background()

		writeTo = os.Stdout
		service = ""
		level   = logger.LevelInfo
	)

	log := logger.New(writeTo, level, service, traceFunc)
	log.Info(ctx, "Startup - Service Details", "logLevel", log.GetLevel().ToString(), "build", build, "cores", runtime.GOMAXPROCS(0))

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "service error, shutting down", "errorDetails", err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	cfg, err := loadConfig(ctx)
	if err != nil {
		return err
	}

	// -----------------------------------------------------------------------
	// init DB

	db, err := sqlx.Open("pgx", "postgresql://pool_ovpq_user:gterb4VES18dEaIyC50VEmzL5twZewmP@dpg-d3himi33fgac739ram8g-a.oregon-postgres.render.com/pool_ovpq")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("db cannot connect: %w", err)
	}

	// // --- Leer archivo SQL ---
	// script, err := ioutil.ReadFile("Reserva-Hotel-BD.sql")
	// if err != nil {
	// 	fmt.Println("no se pudo leer el archivo SQL: %v", err)
	// }

	// // --- Ejecutar el script SQL ---
	// _, err = db.ExecContext(ctx, string(script))
	// if err != nil {
	// 	fmt.Println("error al ejecutar el script SQL: %v", err)
	// }

	// -----------------------------------------------------------------------
	// Repositories

	var (
		usuarioRepository = usuariodb.NewRepository(db)
		reservaRepository = reservadb.NewRepository(db)
		pagoRepository    = pagodb.NewRepository(db)
	)

	// -----------------------------------------------------------------------
	// Services

	var (
		usuarioService = usuario.NewService(usuarioRepository)
		reservaService = reserva.NewService(reservaRepository)
		pagoService    = pago.NewService(pagoRepository, reservaRepository)
	)

	// -----------------------------------------------------------------------
	// Routes

	router := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // Origen de tu frontend (ajusta si es diferente)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Incluye OPTIONS
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Si usas cookies/auth
		MaxAge:           300,  // Tiempo de caché para preflight
	})
	// Aplica el middleware al router
	router.Use(corsHandler.Handler)

	usuario.MakeHandlerWith(usuarioService).SetRoutesTo(router)
	reserva.MakeHandlerWith(reservaService).SetRoutesTo(router)
	pago.MakeHandlerWith(pagoService).SetRoutesTo(router)

	// -------------------------------------------------------------------------
	// HTTP App Server

	var (
		shutdownListener = make(chan os.Signal, 1)
		errListener      = make(chan error, 1)
	)

	signal.Notify(shutdownListener, syscall.SIGINT, syscall.SIGTERM)

	api := http.Server{
		Addr:         "0.0.0.0:4124",
		ReadTimeout:  (time.Second * 15),
		WriteTimeout: (time.Second * 15),
		IdleTimeout:  (time.Second * 15),
		Handler:      router,
	}

	go func() {
		log.Info(ctx, "Startup - API router started", "host", api.Addr)

		errListener <- api.ListenAndServe()
	}()

	// -----------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-errListener:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdownListener:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown completed", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("cannot stop server gracefully: %w", err)
		}
	}

	return nil
}

func traceFunc(ctx context.Context) []any {
	v := tracer.GetValues(ctx)

	fields := make([]any, 2, 4)
	fields[0], fields[1] = "traceID", v.TraceID

	if v.Rut != "" {
		fields = append(fields, "RUT", v.Rut)
	}

	return fields
}
