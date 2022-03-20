package services

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"letshop-backend/config"
	"letshop-backend/models"
)

var Database *gorm.DB

func InitializeDatabase() {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})

	if err != nil {
		panic("[ERROR] Failed to initialize database")
	}

	db.AutoMigrate(
		&models.Profile{},
		&models.Shop{},
		&models.Product{},
		&models.ProductImage{},
	)

	Database = db
}
