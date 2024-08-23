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

// AddEmployee godoc
// @Summary Add an employee
// @Description Add an employee with role
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee"
// @Success 200 {string} string "Employee added successfully."
// @Failure 400 {string} string "Invalid input."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Shop not found."
// @Failure 500 {string} string "Failed to add employee."
// @Router /employee/{shop_id} [post]
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)
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

	if shop.OwnerID != claim.UserID {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var employee models.Employee
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	if result := database.DB.Create(&employee); result.Error != nil {
		http.Error(w, "Failed to add employee.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Employee added successfully."})
	if err != nil {
		return
	}
}

// RemoveEmployee godoc
// @Summary Remove an employee
// @Description Remove an employee from shop
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Employee"
// @Success 200 {string} string "Employee removed successfully."
// @Failure 400 {string} string "The employee is not an employee of this workplace."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Employee not found."
// @Failure 500 {string} string "Failed to remove employee."
// @Router /employee/{shop_id} [delete]
func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.User)
	params := mux.Vars(r)
	shopID, err := strconv.Atoi(params["shop_id"])
	if err != nil {
		http.Error(w, "Invalid shop ID.", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID.", http.StatusBadRequest)
		return
	}

	tx := database.DB.Begin()

	var shop models.Shop
	if result := tx.First(&shop, shopID); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return
	}

	if shop.OwnerID != claim.ID {
		tx.Rollback()
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var employee models.Employee
	if result := tx.First(&employee, "user_id = ?", userID); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Employee not found.", http.StatusNotFound)
		return
	}

	if employee.ShopID != shop.ID {
		tx.Rollback()
		http.Error(w, "The employee is not an employee of this workplace.", http.StatusBadRequest)
		return
	}

	if result := tx.Delete(&employee); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to remove employee.", http.StatusInternalServerError)
		return
	}

	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		http.Error(w, "Database operation could not be completed.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Employee removed successfully."})
	if err != nil {
		return
	}
}

// UpdateEmployeeRole godoc
// @Summary Update an employee
// @Description Update an employee with role
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body models.EmployeeRoleUpdate true "Employee"
// @Success 200 {string} string "Employee role updated successfully."
// @Failure 400 {string} string "Invalid input."
// @Failure 401 {string} string "Unauthorized."
// @Failure 404 {string} string "Employee not found."
// @Failure 500 {string} string "Failed to update employee."
// @Router /employee/{user_id} [put]
func UpdateEmployeeRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID.", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	if result := database.DB.First(&employee, "user_id = ?", userID); result.Error != nil {
		http.Error(w, "Employee not found.", http.StatusNotFound)
		return
	}

	var role models.EmployeeRoleUpdate
	err = json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	employee.Role = role.Role
	employee.UpdatedAt = time.Now()

	if result := database.DB.Save(&employee); result.Error != nil {
		http.Error(w, "Failed to update employee.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Employee role updated successfully."})
	if err != nil {
		return
	}

}
