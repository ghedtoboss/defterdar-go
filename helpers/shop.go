package helpers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"net/http"
)

func FindShopWithClaims(claims *models.Claims, w http.ResponseWriter) (uint, error) {
	var shopID uint
	if result := database.DB.Raw("SELECT id FROM shops WHERE owner_id = ?", claims.UserID).Scan(&shopID); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 0, result.Error
	}
	return shopID, nil

}
