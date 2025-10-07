package reserva

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/web"
)

type Reserva struct {
	IDReserva    int     `json:"IDReserva"`
	IDCliente    int     `json:"IDCliente"`
	IDHabitacion int     `json:"IDHabitacion"`
	FechaInicio  string  `json:"FechaInicio"`
	FechaFin     string  `json:"FechaFin"`
	Total        float64 `json:"Total"`
	Estado       string  `json:"Estado"`
	FechaReserva string  `json:"FechaReserva"`
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
