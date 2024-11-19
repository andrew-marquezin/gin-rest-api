package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.GET("/:name", controllers.Greeting)
	r.Run()
}
