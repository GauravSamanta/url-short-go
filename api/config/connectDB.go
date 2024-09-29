package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {

	dsn := string(os.Getenv("PROD_DB_URL"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{

		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal("Error connecting the database")
	}

}
