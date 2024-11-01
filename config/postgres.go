package config

import (
	"errors"
	"os"

	"github.com/vinicius457/BancoDeDados2/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresql() (*gorm.DB, error) {
	logger := GetLogger(("postgresql"))
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		errDotEnv :=  errors.New("Missing environment variables for database connection")
        logger.Errorf("postgresql connection error: %v", errDotEnv)
		return nil, errDotEnv
    }
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil {
        logger.Errorf("postgresql connection error: %v", err)
        return nil, err
    }

	err = db.AutoMigrate(&model.Cliente{}, &model.Fornecedor{}, &model.Ingrediente{}, &model.Prato{}, &model.Venda{}, &model.Usos{})
	if err!= nil {
        logger.Errorf("postgresql migration error: %v", err)
        return nil, err
    }
	return db, nil
}