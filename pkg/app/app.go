package app

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJSON(req *http.Request, payload any) error {
	if req.Body == nil {
		return errors.New("request body can not be empty")
	}

	err := json.NewDecoder(req.Body).Decode(payload)
	if err != nil {
		return errors.New("invalid request payload")
	}

	return nil
}

func SendErrorResponse(res http.ResponseWriter, status int, errors any) {
	response := map[string]any{
		"success": false,
		"errors":  errors,
	}
	sendJSONResponse(res, status, response)
}

func SendSuccessResponse(res http.ResponseWriter, status int, data any) {
	response := map[string]any{
		"success": true,
		"data":    data,
	}
	sendJSONResponse(res, status, response)
}

func sendJSONResponse(res http.ResponseWriter, status int, data any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}
