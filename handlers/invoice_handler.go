package handlers

import (
	"billing/database"
	"billing/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var trip models.Trip
	if err := database.DB.First(&trip, invoice.TripID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
		return
	}
	database.DB.Create(&invoice)
	c.JSON(http.StatusOK, invoice)

}

func GetInvoice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}
	c.JSON(http.StatusOK, invoice)

}

func GetInvoicesForCustomer(c *gin.Context) {

	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var invoices []models.Invoice
	if err := database.DB.Where("customer_id=?", invoice.CustomerId).Find(&invoices).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}
