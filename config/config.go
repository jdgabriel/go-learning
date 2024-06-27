package config

import "gorm.io/gorm"

// Variável que pode ser usada dentro do package
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error /* return Error ou nada (nil) */ {
	return nil
}

func GetLogger(p string) *Logger {
	// Iniciar Logger
	logger := NewLogger(p)
	return logger
}
