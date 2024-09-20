package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Record a new transaction associated with the authenticated user's shop
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Transaction body"
// @Success 201 {string} string "Transaction created successfully."
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to create transaction"
// @Router /transactions [post]
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusInternalServerError)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var transaction models.Transaction
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusInternalServerError)
		return
	}

	transaction.ShopID = shopID
	if result := database.DB.Create(&transaction); result.Error != nil {
		http.Error(w, "Failed to create transaction.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Transaction created successfully."})
	if err != nil {
		http.Error(w, "Failed to encode transaction.", http.StatusInternalServerError)
		return
	}
}

// UpdateTransaction godoc
// @Summary Update an existing transaction
// @Description Update an existing transaction details associated with the authenticated user's shop
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Param transaction body models.Transaction true "Updated transaction body"
// @Success 200 {string} string "Transaction updated successfully."
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Failed to find transaction"
// @Failure 500 {string} string "Failed to update transaction"
// @Router /transactions/{transaction_id} [put]
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	transactionID, err := strconv.Atoi(params["transaction_id"])
	if err != nil {
		http.Error(w, "Invalid transaction id.", http.StatusInternalServerError)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var transaction models.Transaction
	if result := database.DB.Where("id = ? AND shop_id = ?", transactionID, shopID).First(&transaction); result.Error != nil {
		http.Error(w, "Failed to find transaction.", http.StatusNotFound)
		return
	}

	var updatedTransaction models.Transaction
	err = json.NewDecoder(r.Body).Decode(&updatedTransaction)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusInternalServerError)
		return
	}

	if result := database.DB.Model(&transaction).Updates(updatedTransaction); result.Error != nil {
		http.Error(w, "Failed to update transaction.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Transaction updated successfully."})
	if err != nil {
		http.Error(w, "Failed to update transaction.", http.StatusInternalServerError)
		return
	}
}

// GetTransactions godoc
// @Summary Get list of transactions
// @Description Retrieve a list of transactions for the authenticated user's shop
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to find transactions"
// @Router /transactions [get]
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized.", http.StatusInternalServerError)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var transactions []models.Transaction
	if result := database.DB.Where("shop_id = ?", shopID).Find(&transactions); result.Error != nil {
		http.Error(w, "Failed to find transactions.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, "Failed to encode transactions.", http.StatusInternalServerError)
	}
}

// GetTransaction godoc
// @Summary Get a transaction
// @Description Retrieve a transaction by its ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Success 200 {object} models.Transaction
// @Failure 400 {string} string "Invalid transaction id"
// @Failure 404 {string} string "Failed to find transaction"
// @Failure 500 {string} string "Failed to encode transaction"
// @Router /transactions/{transaction_id} [get]
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID, err := strconv.Atoi(params["transaction_id"])
	if err != nil {
		http.Error(w, "Invalid transaction id.", http.StatusInternalServerError)
		return
	}

	var transaction models.Transaction
	if result := database.DB.Where("id = ?", transactionID).First(&transaction); result.Error != nil {
		http.Error(w, "Failed to find transaction.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, "Failed to encode transaction.", http.StatusInternalServerError)
		return
	}
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete a transaction by its ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Success 200 {string} string "Transaction deleted successfully."
// @Failure 400 {string} string "Invalid transaction id"
// @Failure 404 {string} string "Failed to find transaction"
// @Failure 500 {string} string "Failed to delete transaction"
// @Router /transactions/{transaction_id} [delete]
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID, err := strconv.Atoi(params["transaction_id"])
	if err != nil {
		http.Error(w, "Invalid transaction id.", http.StatusInternalServerError)
		return
	}

	var transaction models.Transaction
	if result := database.DB.Where("id = ?", transactionID).First(&transaction); result.Error != nil {
		http.Error(w, "Failed to find transaction.", http.StatusNotFound)
		return
	}

	if result := database.DB.Delete(&transaction); result.Error != nil {
		http.Error(w, "Failed to delete transaction.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Transaction deleted successfully."})
	if err != nil {
		http.Error(w, "Failed to delete transaction.", http.StatusInternalServerError)
		return
	}
}
