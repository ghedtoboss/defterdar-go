package controllers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateTransaction godoc
// @Summary Create a transaction
// @Description Create a new transaction.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Transaction body models.CustomerTransaction true "Transaction"
// @Success 201 {string} string "Transaction created successfully."
// @Failure 500 {string} string "Failed to create transaction"
// @Failure 400 {string} string "Invalid input."
// @Router /transaction [POST]
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)

	var transaction models.CustomerTransaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	transaction.UserID = claim.UserID

	if result := database.DB.Create(&transaction); result.Error != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Transaction created successfully."})
	if err != nil {
		return
	}
}

// GetTransaction godoc
// @Summary Get a transaction
// @Description Get a transaction.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Transaction body models.CustomerTransaction true "Transaction"
// @Success 201 {object} models.CustomerTransaction
// @Failure 400 {string} string "Failed to create transaction"
// @Failure 404 {string} string "Failed to get transaction."
// @Router /transaction/{transaction_id} [GET]
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID, err := strconv.Atoi(params["transaction_id"])
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	var transaction models.CustomerTransaction
	if result := database.DB.First(&transaction, transactionID); result.Error != nil {
		http.Error(w, "Failed to get transaction", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		return
	}
}
