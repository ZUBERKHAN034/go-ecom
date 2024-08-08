package lib

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJSON(req *http.Request, payload interface{}) error {
	if req.Body == nil {
		return errors.New("request body can not be empty")
	}

	err := json.NewDecoder(req.Body).Decode(payload)
	if err != nil {
		return errors.New("invalid request payload")
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
