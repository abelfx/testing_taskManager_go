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

	// Initialize Mongo collections and repositories
	taskCollection := client.Database("taskdb").Collection("tasks")
	userCollection := client.Database("userdb").Collection("users")

	var taskRepo domain.TaskRepository = repositories.NewMongoTaskRepository(taskCollection)
	var userRepo domain.UserRepository = repositories.NewMongoUserRepository(userCollection)

	// Initialize usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Setup router with all controllers and middleware
	router := routers.SetupRouter(taskController, userController)

	// Start server
	router.Run(":8080")
}
