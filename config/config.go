package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func DBConnection() error {
	var err error
	db, err = ConnectPostgresql()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
        return err
    }
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetDbConnection() *gorm.DB {
	return db
}