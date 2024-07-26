package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// CustomError represents a custom error with a message.
type CustomError struct {
	Message string
}

func (customError *CustomError) Error() string {
	return customError.Message
}

func ParseJSON(req *http.Request, payload any) error {
	if req.Body == nil {
		return &CustomError{Message: "missing request body"}
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return &CustomError{Message: "failed to read request body"}
	}

	if len(body) == 0 {
		return &CustomError{Message: "request body is empty"}
	}

	return json.NewDecoder(req.Body).Decode(payload)
}

func WriteJSON(res http.ResponseWriter, status int, val any) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)

	return json.NewEncoder(res).Encode(val)
}

func WriteError(res http.ResponseWriter, status int, err error) {
	WriteJSON(res, status, map[string]string{"error": err.Error()})
}
