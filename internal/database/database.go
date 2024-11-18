package database

import (
	"fmt"
	"go-chat-system/internal/models" // Adjust import path to your models package
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect establishes a connection to the database
func Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	return nil
}

// Migrate automatically migrates the models (i.e., creates or updates tables)
func Migrate() error {
	if err := DB.AutoMigrate(&models.Chat{}, &models.Message{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
