package handlers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_URI")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to Database", err)
	}

	if err := DB.AutoMigrate(&models.Users{}); err != nil {
		log.Fatalln("Failed to migrate user model to database", err)
	}

	fmt.Println("DB connected Successfully🙌")
}
