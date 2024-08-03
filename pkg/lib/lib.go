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

func SendErrorResponse(res http.ResponseWriter, status int, errors interface{}) {
	response := map[string]interface{}{
		"success": false,
		"errors":  errors,
	}
	sendJSONResponse(res, status, response)
}

func SendSuccessResponse(res http.ResponseWriter, status int, data interface{}) {
	response := map[string]interface{}{
		"success": true,
		"data":    data,
	}
	sendJSONResponse(res, status, response)
}

func sendJSONResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}
