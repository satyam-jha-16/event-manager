package db

import (
	"github.com/satyam-jha-16/event-manager/models"
	"gorm.io/gorm"
)
func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}