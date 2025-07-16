package router

import (
	"restfulapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/tasks", controllers.AddTask)
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)


	return router
}
