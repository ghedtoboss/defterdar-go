package controllers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
)

// CreateShop godoc
// @Summary Create a shop
// @Description Create a shop with name and adress
// @Tags Shop
// @Accept json
// @Produce json
// @Param user body models.Shop true "Shop"
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

	if result := database.DB.Create(&shop); result.Error != nil {
		http.Error(w, "Failed to create shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{
		"messages": "Shop created successfully.",
	})
	if err != nil {
		return
	}
}
