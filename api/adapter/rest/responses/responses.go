package responses

import (
	"fmt"
	"net/http"
)

var errorResponseTemplate = `{
		"errorCode": %v,
		"errorMessage": %v
	}`
var resourceCreatedTemplate = `{
	"statusCode": %v,
	"message": %v
}`

var successTemplate = `{
	"statusCode": %v,
	"data": %v
}`

func BadRequestResponse(w http.ResponseWriter, errorMessage string) {
	payload := fmt.Sprintf(errorResponseTemplate, http.StatusBadRequest, errorMessage)

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(payload))
}

func InternalServerError(w http.ResponseWriter, errorMessage string) {
	payload := fmt.Sprintf(errorResponseTemplate, http.StatusInternalServerError, errorMessage)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(payload))
}

func SuccessOnCreate(w http.ResponseWriter, message string) {
	payload := fmt.Sprintf(resourceCreatedTemplate, http.StatusCreated, message)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(payload))
}

func Success(w http.ResponseWriter, message string) {
	payload := fmt.Sprintf(successTemplate, http.StatusOK, message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(payload))
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
