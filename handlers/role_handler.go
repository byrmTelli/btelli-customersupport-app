package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"encoding/json"
	"net/http"
)

func AssingRoleToUser(w http.ResponseWriter, r *http.Request) {
	var req models.AssignRoleToUserRequestModel

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ApiResponse(w, nil, "Invalid schema.", http.StatusBadRequest)
		return
	}

	//Check is user exist
	var isUserExist models.User
	if err := database.DB.Where("id = ?", req.UserID).First(&isUserExist).Error; err != nil {
		ApiResponse(w, nil, "There is no record mathed given values.", http.StatusBadRequest)
		return
	}

	// Check is role exist
	var isRoleExist models.Role
	if err := database.DB.Where("id = ?", req.RoleID).First(&isRoleExist).Error; err != nil {
		ApiResponse(w, nil, "There is no record mathed given values.", http.StatusBadRequest)
		return
	}

	isUserExist.RoleID = req.RoleID

	if err := database.DB.Save(&isUserExist).Error; err != nil {
		ApiResponse(w, nil, "Error occurred while updating user role.", http.StatusInternalServerError)
		return
	}

	ApiResponse(w, nil, "", http.StatusOK)
}
