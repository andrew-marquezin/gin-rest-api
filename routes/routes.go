package routes

import (
	"gin-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.StudentById)
	r.GET("/:name", controllers.Greeting)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/cpf/:cpf", controllers.StudentByCpf)
	r.Run()
}
