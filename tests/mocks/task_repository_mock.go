package mocks

import (
	"restfulapi/domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) Add(task *domain.Task) (primitive.ObjectID, error) {
	args := m.Called(task)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *TaskRepositoryMock) GetAll() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) GetByID(id primitive.ObjectID) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *TaskRepositoryMock) DeleteByID(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TaskRepositoryMock) UpdateByID(id primitive.ObjectID, task *domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}
