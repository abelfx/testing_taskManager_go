package routers

import (
	"restfulapi/Delivery/controllers"
	infrastructure "restfulapi/infrastructure"
	"restfulapi/usecases"

	"github.com/gin-gonic/gin"
)
func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController, userUsecase *usecases.UserUsecase) *gin.Engine {
	r := gin.Default()

	// Public routes (no auth required)
	r.POST("/signup", userCtrl.Signup)
	r.POST("/login", userCtrl.Login)

	// Routes requiring authentication
	auth := r.Group("/")
	auth.Use(infrastructure.AuthMiddleware())

	// Task routes (authenticated users)
	auth.GET("/tasks", taskCtrl.GetAllTasks)
	auth.GET("/tasks/:id", taskCtrl.GetTask)

	// Admin-only routes
	admin := auth.Group("/")
	admin.Use(infrastructure.AdminOnly(userUsecase))  // <-- pass instance here

	admin.GET("/users", userCtrl.GetUsers)
	admin.POST("/users/promote", userCtrl.PromoteUser)
	admin.DELETE("/users/:id", userCtrl.DeleteUser)

	// Other authenticated task routes
	auth.PUT("/tasks/:id", taskCtrl.UpdateTask)
	auth.DELETE("/tasks/:id", taskCtrl.DeleteTask)
	auth.POST("/tasks", taskCtrl.AddTask)

	return r
}
