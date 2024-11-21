package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/:name", controllers.Greeting)
	r.Run()
}
