package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func ApproveContent(itemId int, w http.ResponseWriter, r *http.Request) bool {

	claimsValue := r.Context().Value("claims")
	if claimsValue == nil {
		ApiResponse(w, nil, "No claim found.", http.StatusUnauthorized)
		return false
	}

	claims, ok := claimsValue.(*jwt.MapClaims)
	if !ok {
		ApiResponse(w, nil, "Invalid claims type.", http.StatusUnauthorized)
		return false
	}

	subject := (*claims)["sub"].(string)
	userID, err := strconv.Atoi(subject)
	if err != nil {
		ApiResponse(w, nil, "Invalid user ID in token.", http.StatusUnauthorized)
		return false
	}

	userRole := (*claims)["role"].(string)
	if (strings.ToLower(userRole) != "admin" || strings.ToLower(userRole) != "help desk") && (itemId != userID) {
		ApiResponse(w, nil, "You have no permission to see this content.", http.StatusUnauthorized)
		return false
	}

	return true
}

func ApiResponse(w http.ResponseWriter, data interface{}, err interface{}, statusCode int) {
	var errorResponse interface{}

	switch e := err.(type) {
	case nil:
		errorResponse = nil
	case string:
		errorResponse = e
	case []string:
		errorResponse = e
	default:
		errorResponse = "An unexpected error occurred"
	}

	response := ApiResponseType{
		Data:  data,
		Error: errorResponse,
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
