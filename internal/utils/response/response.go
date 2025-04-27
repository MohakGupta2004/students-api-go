package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(error validator.ValidationErrors) Response {
	var errors []string
	for _, err := range error {
		switch err.ActualTag() {
		case "required":
			errors = append(errors, fmt.Sprintf("Incorrect format for %s field", err.Field()))
		default:
			errors = append(errors, fmt.Sprintf("Incorrect format for %s field", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errors, ", "),
	}
}
