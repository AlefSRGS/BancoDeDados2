package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() (*gorm.DB ,error) {
	db, err := InitializePostgresql()
	if err != nil {
        return nil, err
    }
	return db, nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}