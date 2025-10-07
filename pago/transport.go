package pago

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpHandler struct {
	svc Service
}

func MakeHandlerWith(svc Service) *HttpHandler {
	return &HttpHandler{svc: svc}
}

func (h *HttpHandler) SetRoutesTo(r chi.Router) {
	r.Post("/crear-pago", h.CreatePayment)
}

func (h *HttpHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	fmt.Printf("%+v", payment)

	if err := h.svc.CreatePayment(r.Context(), payment); err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
	}

	w.WriteHeader(http.StatusNoContent)

}
