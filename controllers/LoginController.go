package controllers

import (
	"github.com/arthuramorim04/go-activity-api.git/database"
	"github.com/arthuramorim04/go-activity-api.git/models"
	"github.com/arthuramorim04/go-activity-api.git/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var loginUser models.Login

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	user.Password = services.SHA256Encode((loginUser.Password))
	user.Email = loginUser.Email
	dbError := db.Where("email = ?", user.Email).Where("password = ?", user.Password).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "User not found credentials",
		})
		return
	}

	token, error := services.NewJWTService().GenerateToken(uint64(user.ID))

	if error != nil {
		c.JSON(400, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
