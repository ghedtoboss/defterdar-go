package controllers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CreateShop godoc
// @Summary Create a shop
// @Description Create a shop with name and adress
// @Tags Shop
// @Accept json
// @Produce json
// @Param shop body models.Shop true "Shop"
// @Success 200 {string} string "Shop created successfully."
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Failed to create shop."
// @Router /shop [post]
func CreateShop(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)

	var shop models.Shop
	err := json.NewDecoder(r.Body).Decode(&shop)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	shop.OwnerID = claim.UserID
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	if result := database.DB.Create(&shop); result.Error != nil {
		http.Error(w, "Failed to create shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Shop created successfully."})
	if err != nil {
		return
	}
}

// UpdateShop godoc
// @Summary Update shop
// @Description Update shop's name and adress
// @Tags Shop
// @Accept json
// @Produce json
// @Param shop body models.Shop true "Shop"
// @Success 200 {string} string "Shop updated successfully."
// @Failure 400 {string} string "Invalid input."
// @Failure 401 {string} string "Unauthorized."
// @Failure 500 {string} string "Failed to update shop."
// @Router /shop/{shop_id} [put]
func UpdateShop(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)
	params := mux.Vars(r)
	shopID, err := strconv.Atoi(params["shop_id"])
	if err != nil {
		http.Error(w, "Invalid shop ID.", http.StatusBadRequest)
		return
	}

	var shop models.Shop
	if result := database.DB.First(&shop, shopID); result.Error != nil {
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return
	}

	if shop.OwnerID != claim.UserID {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var input models.Shop
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	if input.Name != "" {
		shop.Name = input.Name
	}
	if input.Address != "" {
		shop.Address = input.Address
	}

	shop.UpdatedAt = time.Now()

	if result := database.DB.Save(&shop); result.Error != nil {
		http.Error(w, "Failed to update shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Shop updated successfully."})
	if err != nil {
		return
	}

}

// DeleteShop godoc
// @Summary Delete a shop
// @Description Delete a shop
// @Tags Shop
// @Accept json
// @Produce json
// @Param shop body models.Shop true "Shop"
// @Success 200 {string} string "Shop deleted successfully."
// @Failure 400 {string} string "Invalid id."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Shop not found."
// @Failure 500 {string} string "Failed to delete shop."
// @Router /shop/{shop_id} [delete]
func DeleteShop(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(models.Claim)
	params := mux.Vars(r)
	shopID, err := strconv.Atoi(params["shop_id"])
	if err != nil {
		http.Error(w, "Invalid id.", http.StatusBadRequest)
		return
	}

	var shop models.Shop
	if result := database.DB.First(&shop, shopID); result.Error != nil {
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return
	}

	if shop.OwnerID != claims.UserID {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	tx := database.DB.Begin()

	if result := tx.Raw("SET FOREIGN_KEY_CHECKS=0"); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to disabe foreign key.", http.StatusInternalServerError)
		return
	}

	if result := tx.Delete(&shop); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete shop.", http.StatusInternalServerError)
		return
	}

	if result := tx.Raw("SET FOREIGN_KEY_CHECKS=1"); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to enable foreign key.", http.StatusInternalServerError)
		return
	}

	if result := tx.Commit(); result.Error != nil {
		http.Error(w, "Database operation could not be completed.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Shop deleted successfully."})
	if err != nil {
		return
	}
}
