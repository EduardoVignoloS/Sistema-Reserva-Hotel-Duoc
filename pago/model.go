package pago

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/web"
)

type Payment struct {
	IDReserva int     `json:"IDReserva"`
	Monto     float64 `json:"monto"`
	TipoPago  string  `json:"TipoPago"`
	Estado    string  `json:"estado"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func newError(err error) ([]byte, int) {
	var status int
	switch {
	case web.IsRequestError(err):
		errReq := web.GetRequestError(err)
		status = errReq.Status
	default:
		status = http.StatusInternalServerError
	}

	errorResponse := ErrorResponse{
		ErrorMessage: err.Error(),
	}
	errorResponseJSON, err := json.Marshal(errorResponse)
	if err != nil {
		return []byte(err.Error()), http.StatusInternalServerError
	}

	return errorResponseJSON, status
}
