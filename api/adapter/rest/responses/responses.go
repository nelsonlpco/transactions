package responses

import (
	"fmt"
	"net/http"
)

var errorResponseTemplate = `{
		"errorCode": %v,
		"errorMessage": %v
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
