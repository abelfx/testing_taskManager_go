package routers

import (
	"restfulapi/Delivery/controllers"
	"restfulapi/Delivery/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController) *gin.Engine {
	r := gin.Default()

	// Public routes (no auth required)
	r.POST("/signup", userCtrl.Signup)
	r.POST("/login", userCtrl.Login)

	// Routes requiring authentication
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Task routes (authenticated users)
	auth.POST("/tasks", taskCtrl.AddTask)
	auth.GET("/tasks", taskCtrl.GetAllTasks)
	auth.GET("/tasks/:id", taskCtrl.GetTask)
	auth.PUT("/tasks/:id", taskCtrl.UpdateTask)
	auth.DELETE("/tasks/:id", taskCtrl.DeleteTask)

	// Admin-only routes
	admin := auth.Group("/")
	admin.Use(middleware.AdminOnly())

	admin.GET("/users", userCtrl.GetUsers)
	admin.POST("/users/promote", userCtrl.PromoteUser)
	admin.DELETE("/users/:id", userCtrl.DeleteUser)

	return r
}
