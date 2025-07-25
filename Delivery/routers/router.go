package routers

import (
	"restfulapi/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskCtrl *controllers.TaskController) *gin.Engine {
    r := gin.Default()

    r.POST("/tasks", taskCtrl.AddTask)
    r.GET("/tasks", taskCtrl.GetAllTasks)
    r.GET("/tasks/:id", taskCtrl.GetTask)
    r.DELETE("/tasks/:id", taskCtrl.DeleteTask)
    r.PUT("/tasks/:id", taskCtrl.UpdateTask)

    return r
}
