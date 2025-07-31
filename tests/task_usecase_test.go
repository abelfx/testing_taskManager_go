package tests

import (
	"restfulapi/domain"
	"restfulapi/tests/mocks"
	"restfulapi/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	task := &domain.Task{
		Title:       "Learn Go",
		Description: "Write unit tests",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	mockID := primitive.NewObjectID()
	mockRepo.On("Add", task).Return(mockID, nil)

	id, err := usecase.CreateTask(task)

	assert.NoError(t, err)
	assert.Equal(t, mockID, id)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	tasks := []*domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockRepo.On("GetAll").Return(tasks, nil)

	result, err := usecase.GetAllTasks()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id := primitive.NewObjectID()
	task := &domain.Task{ID: id, Title: "Test task"}

	mockRepo.On("GetByID", id).Return(task, nil)

	result, err := usecase.GetTaskByID(id.Hex())

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_InvalidID(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	_, err := usecase.GetTaskByID("invalid-hex")

	assert.Error(t, err)
	assert.Equal(t, "invalid id", err.Error())
}

func TestDeleteTaskByID_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id := primitive.NewObjectID()

	mockRepo.On("DeleteByID", id).Return(nil)

	err := usecase.DeleteTaskByID(id.Hex())

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepositoryMock)
	usecase := usecases.NewTaskUsecase(mockRepo)

	id := primitive.NewObjectID()
	task := &domain.Task{Title: "Updated"}

	mockRepo.On("UpdateByID", id, task).Return(nil)

	err := usecase.UpdateTask(id.Hex(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
