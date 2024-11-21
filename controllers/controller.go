package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"

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

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
