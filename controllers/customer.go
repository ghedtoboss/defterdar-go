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

// AddCustomer godoc
// @Summary Add a new customer
// @Description Create a new customer associated with the authenticated user's shop
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer body"
// @Success 201 {string} string "Customer created successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to create customer"
// @Router /customers [post]
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var customer models.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	customer.ShopID = shopID
	if result := database.DB.Create(&customer); result.Error != nil {
		http.Error(w, "Failed to create customer.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Customer created successfully"})
	if err != nil {
		http.Error(w, "Failed to encode customer.", http.StatusInternalServerError)
		return
	}
}

// UpdateCustomer godoc
// @Summary Update an existing customer
// @Description Update a customer's details for the authenticated user's shop
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer_id path int true "Customer ID"
// @Param customer body models.Customer true "Updated customer body"
// @Success 200 {object} models.Customer
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to update customer"
// @Router /customers/{customer_id} [put]
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["customer_id"])
	if err != nil {
		http.Error(w, "Invalid customer id.", http.StatusBadRequest)
		return
	}
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var customer models.Customer
	if result := database.DB.Where("id = ?", customerID).First(&customer); result.Error != nil {
		http.Error(w, "Failed to find customer.", http.StatusBadRequest)
		return
	}

	if customer.ShopID != shopID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updatedCustomer models.Customer
	err = json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	if result := database.DB.Save(&updatedCustomer); result.Error != nil {
		http.Error(w, "Failed to update customer.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedCustomer)
	if err != nil {
		http.Error(w, "Failed to update customer.", http.StatusInternalServerError)
	}
}

// GetCustomers godoc
// @Summary Get list of customers
// @Description Retrieve a list of customers for the authenticated user's shop
// @Tags Customers
// @Accept json
// @Produce json
// @Success 200 {array} models.Customer
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to find customers"
// @Router /customers [get]
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var customers []models.Customer
	if result := database.DB.Where("shop_id = ?", shopID).Find(&customers); result.Error != nil {
		http.Error(w, "Failed to find customers.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		http.Error(w, "Failed to encode customers.", http.StatusInternalServerError)
		return
	}
}

// GetCustomer godoc
// @Summary Get a customer
// @Description Retrieve a customer's details for the authenticated user's shop
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer_id path int true "Customer ID"
// @Success 200 {object} models.Customer
// @Failure 400 {string} string "Invalid customer id"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to find customer"
// @Router /customers/{customer_id} [get]
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["customer_id"])
	if err != nil {
		http.Error(w, "Invalid customer id.", http.StatusBadRequest)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var customer models.Customer
	if result := database.DB.Where("id = ? AND shop_id = ?", customerID, shopID).First(&customer); result.Error != nil {
		http.Error(w, "Failed to find customer.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, "Failed to encode customer.", http.StatusInternalServerError)
		return
	}
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Delete a customer from the authenticated user's shop
// @Tags Customers
// @Accept json
// @Produce json
// @Param customer_id path int true "Customer ID"
// @Success 200 {string} string "Customer deleted successfully"
// @Failure 400 {string} string "Invalid customer id"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Failed to delete customer"
// @Router /customers/{customer_id} [delete]
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	claims, err := helpers.GetClaims(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	customerID, err := strconv.Atoi(params["customer_id"])
	if err != nil {
		http.Error(w, "Invalid customer id.", http.StatusBadRequest)
		return
	}

	shopID, err := helpers.FindShopWithClaims(claims, w)
	if err != nil {
		http.Error(w, "Failed to find shop id.", http.StatusUnauthorized)
		return
	}

	var customer models.Customer
	if result := database.DB.Where("id = ? AND shop_id = ?", customerID, shopID).First(&customer); result.Error != nil {
		http.Error(w, "Failed to find customer.", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&customer); result.Error != nil {
		http.Error(w, "Failed to delete customer.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
	if err != nil {
		http.Error(w, "Failed to encode customer.", http.StatusInternalServerError)
	}
}
