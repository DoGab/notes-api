package database

import (
	"fmt"

	"github.com/dogab/notes-api/app/model"
	"github.com/dogab/notes-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func ConnectDB() {
	var err error

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	// Connect to the DB and initialize the DB variable
	DBconn, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	err = DBconn.AutoMigrate(&model.Note{})
	if err != nil {
		fmt.Printf("Error migration model Note: %s", err)
	}
	err = DBconn.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Printf("Error migration model User: %s", err)
	}

	fmt.Println("Database Migrated")
}
