package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"btelli-customersupport-app/utils"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequestModel models.LoginRequestModel

	// Decoding Body
	err := json.NewDecoder(r.Body).Decode(&loginRequestModel)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check Username & Password Status
	if loginRequestModel.Username == "" || loginRequestModel.Password == "" {
		http.Error(w, "Username and Password fields are required.", http.StatusBadRequest)
		return
	}

	// Check is user exist
	var user models.User
	result := database.DB.Where("user_name = ?", loginRequestModel.Username).Preload("Role").First(&user)
	if result.Error != nil {
		http.Error(w, "Username or password wrong.", http.StatusUnauthorized)
		return
	}

	// Compare request's password with user's hash.
	if !utils.CheckPasswordHash(loginRequestModel.Password, user.PasswordHash) {
		http.Error(w, "username or password wrong.", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error occured while generating token", http.StatusInternalServerError)
		return
	}

	loginDTO := models.MapUserLoginDTO(user, token)
	ApiResponse(w, loginDTO, "", http.StatusOK)
}
