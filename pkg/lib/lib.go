package lib

import (
	"encoding/json"
	"net/http"
)

// CustomError represents a custom error with a message.
type CustomError struct {
	Message string
}

func (customError *CustomError) Error() string {
	return customError.Message
}

func ParseJSON(req *http.Request, payload interface{}) error {
	if req.Body == nil {
		return &CustomError{Message: "missing request body"}
	}

	err := json.NewDecoder(req.Body).Decode(payload)
	if err != nil {
		return &CustomError{Message: "failed to decode request body"}
	}

	return nil
}

func WriteJSON(res http.ResponseWriter, status int, val any) error {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)

	return json.NewEncoder(res).Encode(val)
}

func WriteError(res http.ResponseWriter, status int, err error) {
	WriteJSON(res, status, map[string]string{"error": err.Error()})
}

func WriteSuccess(res http.ResponseWriter, status int, message string) {
	WriteJSON(res, status, map[string]string{"message": message})
}
