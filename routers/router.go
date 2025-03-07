package routers

import (
	"student-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    studentRoutes := router.Group("/students")
    {
        studentRoutes.GET("/", controllers.GetStudents)
        studentRoutes.GET("/:id", controllers.GetStudentByID)
        studentRoutes.POST("/", controllers.CreateStudent)
        studentRoutes.PUT("/:id", controllers.UpdateStudent)
        studentRoutes.DELETE("/:id", controllers.DeleteStudent)
		studentRoutes.POST("/companies", controllers.CreateCompany)

    }

    return router
}