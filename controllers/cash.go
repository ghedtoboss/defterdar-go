package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
	"time"
)

// CreateCashEntry godoc
// @Summary Create a cash entry
// @Description Create a cash entry with description, amount, type, and optionally link it to a customer
// @Tags Cash
// @Accept json
// @Produce json
// @Param cashentry body models.CashEntry true "CashEntry"
// @Success 201 {string} string "Cash entry created successfully."
// @Failure 400 {string} string "Invalid input."
// @Failure 404 {string} string "Shop or customer not found."
// @Failure 500 {string} string "Failed to create cash entry."
// @Router /cash [post]
func CreateCashEntry(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)
	shop, err := helpers.FindShopByUserID(w, r, int(claim.UserID))
	if err != nil {
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return
	}

	var cashEntry models.CashEntry
	err = json.NewDecoder(r.Body).Decode(&cashEntry)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	cashEntry.UserID = claim.UserID
	cashEntry.ShopID = shop.ID
	cashEntry.CreatedAt = time.Now()
	cashEntry.UpdatedAt = time.Now()

	if cashEntry.CustomerID != nil {
		var customer models.Customer
		if result := database.DB.First(&customer, *cashEntry.CustomerID); result.Error != nil {
			http.Error(w, "Customer not found.", http.StatusNotFound)
			return
		}

		customerTransaction := models.CustomerTransaction{
			UserID:          claim.UserID,
			ShopID:          shop.ID,
			CustomerID:      customer.ID,
			Amount:          cashEntry.Amount,
			TransactionType: cashEntry.EntryType,
			Description:     cashEntry.Description,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if result := database.DB.Create(&customerTransaction); result.Error != nil {
			http.Error(w, "Failed to create customer transaction.", http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&cashEntry); result.Error != nil {
		http.Error(w, "Failed to create cash entry.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Cash entry created successfully."})
	if err != nil {
		return
	}
}
