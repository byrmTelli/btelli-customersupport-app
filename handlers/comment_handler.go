package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateComment(w http.ResponseWriter, r *http.Request) {

	// Get Params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var comment models.Comment

	// Check users rights here...
	//...
	if err := database.DB.First(&comment, id).Error; err != nil {
		ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&comment).Error; err != nil {
		ApiResponse(w, nil, "Error occured while updating item.", http.StatusInternalServerError)
		return
	}

	// Map to DTO
	commentDTO := models.MapCommentToDTO(comment)
	ApiResponse(w, commentDTO, "", http.StatusOK)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {

	// Check user auth.

	// Create a new item
	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		ApiResponse(w, nil, "Error occured while creating new item.", http.StatusInternalServerError)
		return
	}

	// Map item to DTO
	commentDTO := models.MapCommentToDTO(comment)
	ApiResponse(w, commentDTO, "", http.StatusCreated)
}

func RemoveComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&models.Comment{}, id).Error; err != nil {
		ApiResponse(w, nil, "Error occored while deleting related item.", http.StatusInternalServerError)
		return
	}

	ApiResponse(w, nil, "", http.StatusNoContent)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	complaintId, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var complaint models.Complaint
	// Check wheter complaint exist
	isComplaintExist := database.DB.First(&complaint, complaintId)
	if !DatabaseQuerySuccessResult(w, isComplaintExist, complaint) {
		return
	}

	var comments []models.Comment

	// Check users rights here..

	// Get all ıtems from database...
	database.DB.Where("complaint_id = ?", complaintId).Find(&comments)

	// Mapping DTOs
	commentDTOs := models.MapCommentsToDTO(comments)

	// Return all items as a json format.
	ApiResponse(w, commentDTOs, "", http.StatusOK)
}
