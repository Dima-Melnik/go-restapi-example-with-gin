package db

import (
	"log"

	"github.com/Dima-Melnik/go-restapi-example-with-gin/config"
	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	connStr := config.InitDatabaseConn()

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening db connect: %v", err)
	}

	if err = DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed migration models: %v", err)
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Error closing db connect")
	}
}
