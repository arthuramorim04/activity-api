package controllers

import (
	"strconv"

	"github.com/arthuramorim04/go-activity-api.git/database"
	"github.com/arthuramorim04/go-activity-api.git/models"
	"github.com/gin-gonic/gin"
)

func ShowActivity(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var activity models.Activity
	err = db.First(&activity, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find activity: " + err.Error(),
		})
		return
	}

	c.JSON(200, activity)
}

func ShowAllActivity(c *gin.Context) {
	db := database.GetDatabase()

	var activitys []models.Activity
	err := db.Find(&activitys).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list activitys: " + err.Error(),
		})

		return
	}

	c.JSON(200, activitys)
}

func CreateActivity(c *gin.Context) {

	db := database.GetDatabase()

	var activity models.Activity

	err := c.ShouldBindJSON(&activity)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&activity).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create activity: " + err.Error(),
		})

		return
	}

	c.JSON(200, activity)
}

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Activity{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete activity: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditActivity(c *gin.Context) {
	db := database.GetDatabase()

	var activity models.Activity

	err := c.ShouldBindJSON(&activity)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&activity).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create activity: " + err.Error(),
		})
		return
	}

	c.JSON(200, activity)
}
