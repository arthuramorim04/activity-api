package controllers

import (
	"github.com/arthuramorim04/go-books-api.git/database"
	"github.com/arthuramorim04/go-books-api.git/models"
	"github.com/arthuramorim04/go-books-api.git/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find JSON: " + err.Error(),
		})

		return
	}

	user.Password = services.SHA256Encode(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
