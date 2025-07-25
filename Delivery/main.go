package main

import (
	"log"
	"restfulapi/Delivery/controllers"
	"restfulapi/Delivery/routers"
	"restfulapi/domain"
	"restfulapi/infrastructure"
	"restfulapi/repositories"
	"restfulapi/usecases"
)

func main() {
    // Connect to MongoDB
    client, err := infrastructure.ConnectMongo("mongodb://localhost:27017")
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Initialize Mongo collection and repository
    taskCollection := client.Database("taskdb").Collection("tasks")
    var taskRepo domain.TaskRepository = repositories.NewMongoTaskRepository(taskCollection)

    // Initialize usecase/service
    taskUsecase := usecases.NewTaskUsecase(taskRepo)

    // Initialize controller
    taskController := controllers.NewTaskController(taskUsecase)

    // Setup router
    router := routers.SetupRouter(taskController)

    // Start server
    router.Run(":8080")
}
