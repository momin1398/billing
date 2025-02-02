package database

import (
	"billing/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err := gorm.Open(sqlite.Open("billing.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db ", err)

	}
	err = DB.AutoMigrate(&models.User{}, &models.Customer{}, &models.Trip{}, &models.Invoice{})
	if err != nil {

		log.Fatal("failed to mitrate db ", err)
	}
	var count int64
	DB.Model(&models.User{}).Where("role=?", "admin").Count(&count)
	if count == 0 {
		hashedpws, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		DB.Create(&models.User{Email: "admin@example.com", Password: string(hashedpws)})
	}
}
