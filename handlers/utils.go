package handlers

import (
	"encoding/json"
	"net/http"
)

func ApiResponse(w http.ResponseWriter, data interface{}, errMsg string, statusCode int) {

	response := ApiResponseType{
		Data:  data,
		Error: errMsg,
		Code:  statusCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)

}
