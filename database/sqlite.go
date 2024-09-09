package database

import (
	"log"

	"github.com/rounin-rp/email-phisher/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("emailphisher.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Migrate the schema
	migrateSchema(db)

	return db
}

func migrateSchema(db *gorm.DB) {
	// Migrate the schema (automatically create tables based on models)
	err := db.AutoMigrate(&models.Links{}, &models.Email{})
	if err != nil {
		log.Fatal("failed to migrate database schema:", err)
	}

}
