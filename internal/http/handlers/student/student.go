package student

import (
	"encoding/json"
	"errors"
	"github.com/MohakGupta2004/students-api/internal/types"
	"github.com/MohakGupta2004/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validationErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationErr))
			return
		}

		response.WriteJson(w, http.StatusAccepted, map[string]string{
			"success": "OK",
		})
	}
}
