package config

import (
	"fmt"

	"gorm.io/gorm"
)

// Variável que pode ser usada dentro do package
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error /* return Error ou nada (nil) */ {
	var err error

	// Init SQLite database
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error initialize database: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Iniciar Logger
	logger := NewLogger(p)
	return logger
}
