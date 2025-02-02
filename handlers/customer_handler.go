package handlers

import (
	"billing/database"
	"billing/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func GetCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer models.Customer
	if err := database.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}
