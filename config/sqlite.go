package config

import (
	"os"

	"github.com/jdgabriel/go-learning/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("[SQLITE]")
	dbPath := "./database/main.db"

	// Verificação se existe banco de dados criado
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not found, creating...")
		err := os.MkdirAll("./database", os.ModePerm)
		if err != nil {
			logger.Errorf("Failed to create path database: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect database: %v", err)
		return nil, err
	}

	// Migrate database
	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate database: %v", err)
		return nil, err
	}

	return db, nil
}
