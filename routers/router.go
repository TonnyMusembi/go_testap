package routers

import (
	"student-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Removed the group, registered routes directly
	router.GET("/", controllers.GetStudents)
	router.GET("/:id", controllers.GetStudentByID)
	router.POST("/", controllers.CreateStudent)
	router.PUT("/:id", controllers.UpdateStudent)
	router.DELETE("/:id", controllers.DeleteStudent)
	router.POST("/companies", controllers.CreateCompany)
	router.POST("/branches", controllers.CreateBranch)

	return router
}
