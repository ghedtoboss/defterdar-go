package helpers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"net/http"
)

func FindShopByUserID(w http.ResponseWriter, r *http.Request, userID int) (*models.Shop, error) {
	var shopID int
	if result := database.DB.Raw("SELECT shop_id FROM employees WHERE user_id = ?", userID).Scan(&shopID); result.Error != nil {
		http.Error(w, "Shop id not found.", http.StatusNotFound)
		return nil, result.Error
	}

	var shop models.Shop
	if result := database.DB.First(&shop, shopID); result.Error != nil {
		http.Error(w, "Shop not found.", http.StatusNotFound)
		return nil, result.Error
	}

	return &shop, nil
}
