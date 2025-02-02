package handlers

import (
	"billing/database"
	"billing/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTrip(c *gin.Context) {
	var trip models.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&trip)
	c.JSON(http.StatusOK, trip)

}

func GetTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var trip models.Trip
	if err := database.DB.First(&trip, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
		return
	}
	c.JSON(http.StatusOK, trip)

}
