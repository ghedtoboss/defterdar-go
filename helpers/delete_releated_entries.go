package helpers

import "gorm.io/gorm"

func DeleteReleatedEntries(tx *gorm.DB, table string, shopID uint) error {
	return tx.Exec("DELETE FROM "+table+" WHERE shop_id = ?", shopID).Error
}
