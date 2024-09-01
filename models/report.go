package models

type ReportRequest struct {
	ReportType string `json:"report_type"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
}

type EmployeeReportRequest struct {
	EmployeeID int    `json:"employee_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
}
