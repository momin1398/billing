package database

import (
	"billing/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func Init() {
	var err error
	// Use modernc.org/sqlite (CGO-free SQLite driver)
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
	// AutoMigrate models
	err1 := DB.AutoMigrate(&models.User{}, &models.Customer{}, &models.Trip{}, &models.Invoice{})
	if err1.Error != nil {
		log.Fatal("failed to migrate db ", err)
	}

	// Ensure an admin user exists
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		DB.Create(&models.User{Email: "admin12@example.com", Password: string(hashedPwd), Role: "admin"})
		fmt.Println("line 36")
	}
}
