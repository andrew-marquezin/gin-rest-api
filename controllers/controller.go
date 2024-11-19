package controllers

import (
	"gin-rest-api/models"

	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hi " + name + "! What's up?",
	})
}
