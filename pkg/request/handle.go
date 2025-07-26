package request

import (
	"API-email-confirmation/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {
	defer req.Body.Close()

	body, err := Decode[T](req.Body)
	if err != nil {
		response.JSONResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.JSONResponse(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, err
}
