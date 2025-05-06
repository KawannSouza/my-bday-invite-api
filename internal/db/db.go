package db

import (
	"fmt"
	"log"
	"os"

	"github.com/KawannSouza/my-bday-invite-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db 
	fmt.Println("✅ Database connection established!")
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{}, &model.Invite{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Table migration complete!")
}