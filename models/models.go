package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"Phone"`
	Trips []Trip
}
type Trip struct {
	gorm.Model
	CustomerID  uint    `json:"customer_id"`
	Destination string  `json:"destination"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

type Invoice struct {
	gorm.Model
	CustomerId uint    `json:"customer_id"`
	TripID     uint    `json:"trip_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
