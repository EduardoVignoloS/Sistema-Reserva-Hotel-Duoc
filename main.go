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
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario"
	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/usuario/repository/usuariodb"
	"github.com/go-chi/chi/v5"
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

	db, err := sqlx.Open("pgx", "postgresql://postgres:reserva@db.pktpivqgghdfgtwmnske.supabase.co:5432/postgres?pgbouncer=true&pool_mode=transaction")
	if err != nil {
		return err
	}
	// -----------------------------------------------------------------------
	// Repositories

	var (
		usuarioRepository = usuariodb.NewRepository(db)
	)

	// -----------------------------------------------------------------------
	// Services

	var (
		usuarioService = usuario.NewService(usuarioRepository)
	)

	// -----------------------------------------------------------------------
	// Routes

	router := chi.NewRouter()
	usuario.MakeHandlerWith(usuarioService).SetRoutesTo(router)

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
