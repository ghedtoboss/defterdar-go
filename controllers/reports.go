package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
	"time"
)

// GetCashReports godoc
// @Summary Get report
// @Description Get reports of shop or user
// @Tags Report
// @Accept json
// @Produce json
// @Param report body models.ReportRequest true "Report"
// @Success 200 {object} models.CashEntry
// @Failure 400 {string} string "Invalid date range."
// @Failure 404 {string} string "Shop not found."
// @Failure 500 {string} string "Failed to encode data."
// @Router /report [post]
func GetCashReports(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user").(models.Claim)

	var request models.ReportRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	fromDate, err := time.Parse("2006-01-02", request.FromDate)
	toDate, err := time.Parse("2006-01-02", request.ToDate)
	if err != nil || fromDate.After(toDate) {
		http.Error(w, "Invalid date range.", http.StatusBadRequest)
		return
	}

	var cashEntries []models.CashEntry

	switch request.ReportType {
	case "shop":
		shop, err := helpers.FindShopByUserID(w, r, int(claim.UserID))
		if err != nil {
			http.Error(w, "Shop not found.", http.StatusNotFound)
			return
		}

		result := database.DB.Where("shop_id AND created_at BETWEEN ? AND ?", shop.ID, request.FromDate, request.ToDate).Find(&cashEntries)
		if result.Error != nil {
			http.Error(w, "Failed to get cash entries for shop.", http.StatusInternalServerError)
			return
		}

	case "user":
		result := database.DB.Where("user_id = ? AND created_at BETWEEN ? AND ?", claim.UserID, request.FromDate, request.ToDate).Find(&cashEntries)
		if result.Error != nil {
			http.Error(w, "Failed to get cash entries for user.", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Invalid report type.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cashEntries)
	if err != nil {
		http.Error(w, "Failed to encode data.", http.StatusInternalServerError)
		return
	}
}

// GenerateEmployeeCashReport godoc
// @Summary Generate a cash report for an employee
// @Description Generate a cash report for a specific employee within a shop owned by the user
// @Tags Report
// @Accept json
// @Produce json
// @Param report_request body models.EmployeeReportRequest true "Report Request"
// @Success 200 {array} models.CashEntry
// @Failure 400 {string} string "Invalid input."
// @Failure 403 {string} string "Forbidden: Only owners can access this endpoint."
// @Failure 404 {string} string "Shop or employee not found."
// @Failure 500 {string} string "Failed to generate report."
// @Router /report/employee [post]
func GetEmployeeCashReport(w http.ResponseWriter, r *http.Request) {
	claim, ok := r.Context().Value("user").(models.Claim)
	if !ok {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var request models.EmployeeReportRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	fromDate, err := time.Parse("2006-01-02", request.FromDate)
	if err != nil {
		http.Error(w, "Invalid date format.", http.StatusBadRequest)
		return
	}

	toDate, err := time.Parse("2006-01-02", request.ToDate)
	if err != nil {
		http.Error(w, "Invalid date format.", http.StatusBadRequest)
		return
	}

	if fromDate.After(toDate) {
		http.Error(w, "Invalid date range.", http.StatusBadRequest)
		return
	}

	shop, err := helpers.FindShopByUserID(w, r, int(claim.UserID))
	if err != nil {
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return
	}

	var employee models.Employee
	if result := database.DB.First(&employee, "id = ? AND shop_id = ?", request.EmployeeID, shop.ID); result.Error != nil {
		http.Error(w, "Employee not found or does not belong to your shop.", http.StatusNotFound)
		return
	}

	var cashEntries []models.CashEntry

	if result := database.DB.Where("user_id = ? AND created_at BETWEEN ? AND ?",
		employee.UserID, fromDate, toDate).Find(&cashEntries); result.Error != nil {
		http.Error(w, "Failed to get cash entries.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cashEntries)
	if err != nil {
		http.Error(w, "Failed to encode report data.", http.StatusInternalServerError)
		return
	}
}
