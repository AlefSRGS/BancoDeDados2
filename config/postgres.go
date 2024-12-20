package config

import (
	"os"

	"github.com/vinicius457/BancoDeDados2/internal/model"
	"github.com/vinicius457/BancoDeDados2/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
    host = "postgres_db"
    port = "5432"
    user = "postgres"
    password = "postgres"
    dbname = "restaurante_db"
)

func ConnectPostgresql() (*gorm.DB, error) {
	// logger := GetLogger(("postgresql"))
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	// if host == "" || port == "" || user == "" || password == "" || dbname == "" {
	// 	errDotEnv :=  errors.New("Missing environment variables for database connection")
    //     logger.Warnf("postgresql connection warning: %v", errDotEnv)
    //     logger.Info("Initializing default database connection...")
        // host := "postgres_db"
        // port := "5432"
        // user := "postgres"
        // password := "postgres"
        // dbname := "restaurante"
    //}
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

    err = initDatabase(db)
    if err!= nil {
        logger.Warnf("Warning... something wrong with initializing database: %v", err)
        //return nil, err
    }

	return db, nil
}

func initDatabase(db *gorm.DB) error {
    var sqlFiles []string
    //sqlFiles = append(sqlFiles,"internal/migrations/create_tables.sql")
    sqlFiles = append(sqlFiles,"internal/migrations/create_functions.sql")
    sqlFiles = append(sqlFiles,"internal/migrations/create_procedures.sql")
    sqlFiles = append(sqlFiles,"internal/migrations/create_views.sql")
    sqlFiles = append(sqlFiles,"internal/migrations/create_users_db.sql")
    sqlFiles = append(sqlFiles,"internal/migrations/create_triggers.sql")
    //sqlFiles = append(sqlFiles,"internal/migrations/inserts_tables.sql")

    for _, sqlFile := range sqlFiles {
        err := execSqlScripts(sqlFile, db)
        if err != nil {
            return err
        }
    }

    err := repository.InsertClientes(db)
    if err != nil {
        logger.Warnf("Error with insert in cliente table: %v", err)
    }
    err = repository.InsertFornecedores(db)
    if err != nil {
        logger.Warnf("Error with insert in fornecedor table: %v", err)
    }
    err = repository.InsertIngredientes(db)
    if err != nil {
        logger.Warnf("Error with insert in ingrediente table: %v", err)
    }
    err = repository.InsertPratos(db)
    if err != nil {
        logger.Warnf("Error with insert in prato table: %v", err)
    }
    err = repository.InsertVendas(db)
    if err != nil {
        logger.Warnf("Error with insert in venda table: %v", err)
    }
    return nil
}

func execSqlScripts(sqlFile string, db *gorm.DB) error{
    sqlBytes, err := os.ReadFile(sqlFile)
    if err != nil {
        logger.Errorf("Error reading SQL file for init database: %v", err)
        return err
    }

    commands := string(sqlBytes)
    if result := db.Exec(commands); result.Error != nil {
        logger.Errorf("Error executing SQL command: %v",result.Error)
        return result.Error
    }
    return nil
}