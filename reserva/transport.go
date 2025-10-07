package reserva

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// A handler is similar to a Controller.
type HttpHandler struct {
	svc Service
}

func MakeHandlerWith(svc Service) *HttpHandler {
	return &HttpHandler{svc: svc}
}

func (h *HttpHandler) SetRoutesTo(r chi.Router) {

	r.Post("/crear-reserva", h.createReserva)
	r.Get("/user/reservas", h.reservas)
	r.Put("/reserva/{id}", h.updateReserva)
	r.Delete("/reserva/{id}", h.deleteReserva)

}

func (h *HttpHandler) createReserva(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reserva Reserva
	if err := json.NewDecoder(r.Body).Decode(&reserva); err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	if err := h.svc.CreateReservation(r.Context(), reserva); err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HttpHandler) reservas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ID := r.URL.Query().Get("id")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID es obligatorio"))
		return
	}

	fmt.Println(ID, "ID recibido")
	IDStr, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID debe ser un número entero"))
		return
	}

	reserva, err := h.svc.GetReservationByID(r.Context(), int(IDStr))
	if err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
		return
	}

	response, err := json.Marshal(reserva)
	if err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errJSON)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)

}

func (h *HttpHandler) updateReserva(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reserva Reserva
	if err := json.NewDecoder(r.Body).Decode(&reserva); err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	if err := h.svc.UpdateReservation(r.Context(), reserva); err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) deleteReserva(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID es obligatorio"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("El ID debe ser un número entero"))
		return
	}

	if err := h.svc.DeleteReservation(r.Context(), id); err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
