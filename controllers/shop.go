package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
	"time"
)

// CreateShop godoc
// @Summary Create a new shop
// @Description Create a new shop for the user
// @Tags shop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shop body models.Shop true "Shop"
// @Success 201 {object} models.Shop
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to create shop"
// @Router /shop [post]
func CreateShop(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var count int64
	if result := database.DBRead.Model(&models.Shop{}).Where("owner_id = ?", claims.UserID).Count(&count); result.Error != nil {
		http.Error(w, "kontrol sağlanamadı.", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Zaten bir iş yeriniz var.", http.StatusBadRequest)
		return
	}

	var shop models.Shop
	err = json.NewDecoder(r.Body).Decode(&shop)
	if err != nil {
		http.Error(w, "Invalid request body.", http.StatusBadRequest)
		return
	}

	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()
	shop.OwnerID = claims.UserID

	if result := database.DBWrite.Create(&shop); result.Error != nil {
		http.Error(w, "Failed to create shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(shop)
	if err != nil {
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError)
		return
	}
}

// GetShop godoc
// @Summary Get shop details
// @Description Get details of the shop for the authenticated user
// @Tags shop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.Shop
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to get shop"
// @Router /shop [get]
func GetShop(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var shop models.Shop
	if result := database.DBRead.Preload("Owner").Where("owner_id = ?", claims.UserID).First(&shop); result.Error != nil {
		http.Error(w, "Failed to get shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(shop)
	if err != nil {
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError)
		return
	}
}

// UpdateShop godoc
// @Summary Update shop details
// @Description Update the details of the shop for the authenticated user
// @Tags shop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param shop body models.Shop true "Shop"
// @Success 200 {object} models.Shop
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to update shop"
// @Router /shop [put]
func UpdateShop(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var shop models.Shop
	if result := database.DBRead.Preload("Owner").Where("owner_id = ?", claims.UserID).First(&shop); result.Error != nil {
		http.Error(w, "Failed to get shop.", http.StatusInternalServerError)
		return
	}

	if shop.OwnerID != claims.UserID {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var shopUpdate models.Shop
	err = json.NewDecoder(r.Body).Decode(&shopUpdate)
	if err != nil {
		http.Error(w, "Invalid request body.", http.StatusBadRequest)
		return
	}

	shop.Name = shopUpdate.Name
	shop.Address = shopUpdate.Address
	shop.Phone = shopUpdate.Phone
	shop.UpdatedAt = time.Now()

	if result := database.DBWrite.Save(&shop); result.Error != nil {
		http.Error(w, "Failed to update shop.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(shop)
	if err != nil {
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError)
		return
	}
}

// DeleteShop godoc
// @Summary Delete shop
// @Description Delete the shop and all related records for the authenticated user
// @Tags shop
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} string "Shop deleted successfully"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to delete shop"
// @Router /shop [delete]
func DeleteShop(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var shop models.Shop
	if result := database.DBRead.First(&shop, "owner_id = ?", claims.UserID); result.Error != nil {
		http.Error(w, "Failed to get shop.", http.StatusInternalServerError)
		return
	}
	tx := database.DBWrite.Begin()

	tables := []string{
		"cash_entries",
		"customer_transaction",
		"customers",
		"employees",
		"expenses",
		"incomes",
		"invoices",
		"ledgers",
		"products",
		"receipts",
		"requests",
		"sales",
	}

	for _, table := range tables {
		if err := helpers.DeleteReleatedEntries(tx, table, shop.ID); err != nil {
			tx.Rollback()
			http.Error(w, "Failed to delete entries from"+table, http.StatusInternalServerError)
			return
		}
	}

	if result := tx.Delete(&shop).Error; result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete shop.", http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Shop deleted successfully."})
	if err != nil {
		http.Error(w, "Failed to encode response.", http.StatusInternalServerError)
	}

}
