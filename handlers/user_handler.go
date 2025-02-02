package handlers

import (
	"billing/database"
	"billing/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only admin can create user"})

		return
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hashed"})
		return
	}
	user.Password = string(hashedPwd)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
