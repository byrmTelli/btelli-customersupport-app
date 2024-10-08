package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"btelli-customersupport-app/utils"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequestModel

	// Validate request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ApiResponse(w, nil, "Invalid schema.", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequestModel(req)
	if len(errors) > 0 {
		errorMessages := make([]string, len(errors))
		for i, err := range errors {
			errorMessages[i] = err.Message
		}
		ApiResponse(w, nil, errorMessages, http.StatusBadRequest)
		return
	}

	if req.Password != req.PasswordConfirm {
		ApiResponse(w, nil, "Passwords did not match.", http.StatusBadRequest)
		return
	}

	// Trim spaces maded by mistake
	utils.TrimSpacesInStruct(&req)

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		ApiResponse(w, nil, "Error occurred.", http.StatusInternalServerError)
		return
	}

	// Check if user exists
	var isUserExist models.User
	if err := database.DB.Where("user_name = ? OR email = ? OR phone = ?", req.Username, req.Email, req.Phone).First(&isUserExist).Error; err == nil {
		ApiResponse(w, nil, "This user already exists.", http.StatusBadRequest)
		return
	}

	user := models.User{
		UserName:     req.Username,
		Name:         req.Name,
		Surname:      req.Surname,
		Email:        req.Email,
		Phone:        "0 " + req.Phone,
		PasswordHash: pass,
		RoleID:       3,
	}

	// Create the user and check for errors
	if err := database.DB.Create(&user).Error; err != nil {
		ApiResponse(w, nil, "Error occurred while creating user.", http.StatusInternalServerError)
		return
	}

	ApiResponse(w, nil, "", http.StatusCreated)
}
