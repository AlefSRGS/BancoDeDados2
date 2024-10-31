package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func SetupDataBase() (*gorm.DB, error) {
	logger := GetLogger(("postgresql"))
	//gorm.Open(postgres.Open(), )
}