package usuario

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoVignoloS/Sistema-Reserva-Hotel-Duoc/go-ms-reserva-hotel/kit/web"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type Usuario struct {
	ID             string `json:"ID"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Telefono       string `json:"telefono"`
	TypeC          string `json:"typeC"`
	Fecha_Registro string `json:"fechaRegistro"`
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
