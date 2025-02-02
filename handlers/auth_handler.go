package handlers

import (
	"billing/database"
	"billing/middleware"
	"billing/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginDetails models.User
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := database.DB.Where("email=?", loginDetails.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credential"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credential"})
		return

	}
	token, err := middleware.CreateToken(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not able to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
