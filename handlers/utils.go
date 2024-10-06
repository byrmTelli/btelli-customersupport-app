package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
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

func DatabaseQuerySuccessResult(w http.ResponseWriter, result *gorm.DB, data interface{}) bool {

	switch result.Error {
	case nil:
		return true
	case gorm.ErrRecordNotFound:
		ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		return false
	default:
		ApiResponse(w, nil, "Error occurred while retrieving item.", http.StatusInternalServerError)
		return false
	}
}
