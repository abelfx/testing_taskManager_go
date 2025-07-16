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

	val := data.Add(&task) 
	
	if val {
		ctx.JSON(http.StatusCreated, task)
		return 
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "task with this ID already exists"})
	
	

}

func GetAllTasks(ctx *gin.Context) {
    tasks := data.GetAll()

	if len(tasks) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No Task in DB"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, val := data.Get(id)

	if val {
		ctx.JSON(http.StatusOK,task)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	
	


}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	val := data.Delete(id)

	if val {
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

	success := data.Update(id, &updatedTask)
	if success {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
}
