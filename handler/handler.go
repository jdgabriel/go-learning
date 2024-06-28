package handler

import (
	"github.com/jdgabriel/go-learning/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandlers() {
	logger = config.GetLogger("[HANDLER]")
	db = config.GetSQLite()
}
