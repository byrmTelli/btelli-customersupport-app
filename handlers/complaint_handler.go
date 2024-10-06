package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UpdateComplaint(w http.ResponseWriter, r *http.Request) {

	// Get Params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter.", http.StatusBadRequest)
		return
	}

	var complaint models.Complaint
	if err := database.DB.First(&complaint, id).Error; err != nil {
		ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&complaint)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&complaint).Error; err != nil {
		ApiResponse(w, nil, "Error occured while updating item.", http.StatusInternalServerError)
		return
	}

	// Map to DTO
	complaintDTO := models.MapComplaintToDTO(complaint)
	ApiResponse(w, complaintDTO, "", http.StatusOK)
}

func CreateComplaint(w http.ResponseWriter, r *http.Request) {

	// Check user auth.

	// Create a new item
	var complaint models.Complaint

	err := json.NewDecoder(r.Body).Decode(&complaint)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&complaint).Error; err != nil {
		ApiResponse(w, nil, "Error occured while creating new item.", http.StatusInternalServerError)
		return
	}

	// Map item to DTO
	complaintDTO := models.MapComplaintToDTO(complaint)
	ApiResponse(w, complaintDTO, "", http.StatusCreated)
}

func RemoveComplaint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&models.Complaint{}, id).Error; err != nil {
		ApiResponse(w, nil, "Error occored while deleting related item.", http.StatusInternalServerError)
		return
	}

	ApiResponse(w, nil, "", http.StatusNoContent)
}

func GetComplaint(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		ApiResponse(w, nil, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var complaint models.Complaint

	// Check users rights here...

	// Get related item from database
	result := database.DB.First(&complaint, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		} else {
			ApiResponse(w, nil, "Error occured while retriewing ıtem.", http.StatusInternalServerError)
		}
		return
	}

	// Mapping DTO

	complaintDTO := models.MapComplaintToDTO(complaint)
	// Return related item as a json format.
	ApiResponse(w, complaintDTO, "", http.StatusOK)
	end := time.Now()
	fmt.Println("Time: ", end.Sub(start))
}

func GetComplaints(w http.ResponseWriter, r *http.Request) {
	var complaints []models.Complaint

	// Check users rights here..

	// Get all ıtems from database...
	if err := database.DB.Find(&complaints).Error; err != nil {
		ApiResponse(w, nil, "An error occured while fetching data from database.", http.StatusBadRequest)
		return
	}

	// Mapping DTOs
	complaintDTOs := models.MapComplaintsToDTO(complaints)

	// Return all items as a json format.
	ApiResponse(w, complaintDTOs, "", http.StatusOK)
}
