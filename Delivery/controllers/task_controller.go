package controllers

import (
	"net/http"

	"restfulapi/domain"
	usecases "restfulapi/usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
    taskUsecase *usecases.TaskUsecase
}

func NewTaskController(u *usecases.TaskUsecase) *TaskController {
    return &TaskController{taskUsecase: u}
}

func (ctrl *TaskController) AddTask(c *gin.Context) {
    var task domain.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := ctrl.taskUsecase.CreateTask(&task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add task"})
        return
    }
    task.ID = id
    c.JSON(http.StatusCreated, task)
}

func (ctrl *TaskController) GetAllTasks(c *gin.Context) {
    tasks, err := ctrl.taskUsecase.GetAllTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch tasks"})
        return
    }
    if len(tasks) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) GetTask(c *gin.Context) {
    id := c.Param("id")
    task, err := ctrl.taskUsecase.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func (ctrl *TaskController) DeleteTask(c *gin.Context) {
    id := c.Param("id")
    err := ctrl.taskUsecase.DeleteTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (ctrl *TaskController) UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var updatedTask domain.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := ctrl.taskUsecase.UpdateTask(id, &updatedTask)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}
