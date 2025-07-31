package tests

import (
	"errors"
	"testing"

	"restfulapi/domain"
	"restfulapi/tests/mocks"
	"restfulapi/usecases"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	user := &domain.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	id := primitive.NewObjectID()
	mockRepo.On("Create", user).Return(id, nil)

	returnedID, err := userUsecase.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, id, returnedID)

	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	username := "testuser"
	password := "password123"
	user := &domain.User{
		Username: username,
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	mockRepo.On("Authenticate", username, password).Return(user, nil)

	returnedUser, err := userUsecase.Authenticate(username, password)
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_Failure(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	username := "wronguser"
	password := "wrongpassword"

	mockRepo.On("Authenticate", username, password).Return(nil, errors.New("invalid credentials"))

	returnedUser, err := userUsecase.Authenticate(username, password)
	assert.Error(t, err)
	assert.Nil(t, returnedUser)

	mockRepo.AssertExpectations(t)
}

func TestPromoteUser_InvalidRole(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	userID := primitive.NewObjectID()

	err := userUsecase.PromoteUser(userID, "")
	assert.Error(t, err)
	assert.Equal(t, "new role must be provided", err.Error())
}

func TestPromoteUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	userID := primitive.NewObjectID()
	newRole := "admin"

	mockRepo.On("PromoteUser", userID, newRole).Return(nil)

	err := userUsecase.PromoteUser(userID, newRole)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUserByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	userID := primitive.NewObjectID()

	mockRepo.On("DeleteByID", userID).Return(nil)

	err := userUsecase.DeleteUserByID(userID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUserByID_Failure(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userUsecase := usecases.NewUserUsecase(mockRepo)

	userID := primitive.NewObjectID()

	mockRepo.On("DeleteByID", userID).Return(errors.New("user not found"))

	err := userUsecase.DeleteUserByID(userID)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	mockRepo.AssertExpectations(t)
}
