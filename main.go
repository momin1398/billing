package main

import (
	"billing/database"
	"billing/handlers"
	"billing/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.POST("/login", handlers.Login)
	r.POST("/register", middleware.AuthMiddleware(), middleware.AdminOnly(), handlers.CreateUser)
	r.POST("/customers", middleware.AuthMiddleware(), handlers.CreateCustomer)
	r.GET("/customers/:id", middleware.AuthMiddleware(), handlers.GetCustomer)

	r.POST("/trips", middleware.AuthMiddleware(), handlers.CreateTrip)
	r.GET("/trips/:id", middleware.AuthMiddleware(), handlers.GetTrip)

	r.POST("/invoices", middleware.AuthMiddleware(), handlers.CreateInvoice)
	r.GET("/invoices/:id", middleware.AuthMiddleware(), handlers.GetInvoice)
	r.GET("/customers/:id/invoices", middleware.AuthMiddleware(), handlers.GetInvoicesForCustomer)

	r.Run(":8080")

}
