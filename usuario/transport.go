package usuario

import (
	"encoding/json"
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

	r.Post("/create", h.createUser)
	r.Post("/login", h.login)

}

func (h *HttpHandler) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usuario Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	if err := h.svc.CreateAccount(r.Context(), usuario); err != nil {
		errJSON, status := newError(err)
		w.WriteHeader(status)
		w.Write(errJSON)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

func (h *HttpHandler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usuario Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	usuario, err := h.svc.Login(r.Context(), usuario)
	if err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJSON)
		return
	}

	response, err := json.Marshal(usuario)
	if err != nil {
		errJSON, _ := newError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errJSON)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)

}
