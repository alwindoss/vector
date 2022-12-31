package storage

import "gorm.io/gorm"

func Setup(db *gorm.DB) error {
	db.AutoMigrate(&Station{})
	return nil
}
