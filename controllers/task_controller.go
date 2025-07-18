package controllers

import (
	"net/http"
	"restfulapi/data"
	"restfulapi/models"

	"github.com/gin-gonic/gin"
)

func AddTask(ctx *gin.Context){
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err:= data.Add(&task) 
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError , gin.H{"message": "Failed to add Task to the database"})
		return 
	}

	task.ID = id
	ctx.JSON(http.StatusCreated, task)
}

func GetAllTasks(ctx *gin.Context) {
    tasks, err := data.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch tasks"})
		return
	}

	if len(tasks) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.Get(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK,task)
	
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := data.Delete(id)

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted Successfully"} )
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.Update(id, &updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	} else {

		ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	}
}
