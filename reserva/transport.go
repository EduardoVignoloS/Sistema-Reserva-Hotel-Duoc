package reserva

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// A handler is similar to a Controller.
type HttpHandler struct {
	svc Service
}

func MakeHandlerWith(svc Service) *HttpHandler {
	return &HttpHandler{svc: svc}
}

func (h *HttpHandler) SetRoutesTo(r chi.Router) {
	r.Use(middleware.Logger)

	r.Post("/crear-reserva", h.createReserva)
	r.Get("/reservas", h.reservas)
	r.Put("/reserva/{id}", h.updateReserva)
	r.Delete("/reserva/{id}", h.deleteReserva)

}

func (h *HttpHandler) createReserva(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crear Reserva!"))
}

func (h *HttpHandler) reservas(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar Reservas!"))
}

func (h *HttpHandler) updateReserva(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Actualizar Reserva!"))
}

func (h *HttpHandler) deleteReserva(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eliminar Reserva!"))
}
