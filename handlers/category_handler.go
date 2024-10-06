package handlers

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	// Get Params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter.", http.StatusBadRequest)
		return
	}

	var category models.ComplaintCategory
	if err := database.DB.First(&category, id).Error; err != nil {
		ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&category).Error; err != nil {
		ApiResponse(w, nil, "Error occured while updating item.", http.StatusInternalServerError)
		return
	}

	// Map to DTO
	categoryDTO := models.MapCategoryToDTO(category)
	ApiResponse(w, categoryDTO, "", http.StatusOK)
}
func CreateCategory(w http.ResponseWriter, r *http.Request) {

	// Check user auth.

	// Create a new item
	var category models.ComplaintCategory

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		ApiResponse(w, nil, "Invalid input.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&category).Error; err != nil {
		ApiResponse(w, nil, "Error occured while creating new item.", http.StatusInternalServerError)
		return
	}

	// Map item to DTO
	categoryDTO := models.MapCategoryToDTO(category)
	ApiResponse(w, categoryDTO, "", http.StatusCreated)
}
func RemoveCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ApiResponse(w, nil, "Invalid parameter.", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&models.ComplaintCategory{}, id).Error; err != nil {
		ApiResponse(w, nil, "Error occored while deleting related item.", http.StatusInternalServerError)
		return
	}

	ApiResponse(w, nil, "", http.StatusNoContent)
}
func GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		ApiResponse(w, nil, "Invalid parameter", http.StatusBadRequest)
		return
	}

	var category models.ComplaintCategory

	// Check users rights here...

	result := database.DB.First(&category, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ApiResponse(w, nil, "Item not found.", http.StatusNotFound)
		} else {
			ApiResponse(w, nil, "Error occured while retriewing ıtem.", http.StatusInternalServerError)
		}
		return
	}

	// Mapping
	categoryDTO := models.MapCategoryToDTO(category)
	ApiResponse(w, categoryDTO, "", http.StatusOK)
}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.ComplaintCategory

	// Check users rights here...

	if err := database.DB.Find(&categories).Error; err != nil {
		ApiResponse(w, nil, "An error occured while fetching data from database.", http.StatusBadRequest)
		return
	}

	// Mapping

	categoryDTOs := models.MapCategoriesToDTO(categories)

	ApiResponse(w, categoryDTOs, "", http.StatusOK)
}
