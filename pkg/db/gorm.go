package db

import (
	"SendEmail-Service/pkg/config"
	"SendEmail-Service/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var dsn string
	if len(config.Config("DBUrl")) == 0 {
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Config("DB_HOST"),
			5432,
			config.Config("DB_USER"),
			config.Config("DB_NAME"),
			config.Config("DB_PASS"))
	} else {
		dsn = config.Config("DBUrl")
	}
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
		panic(err)

	}
	//initialize auto migration to create tables
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return database, nil
}
