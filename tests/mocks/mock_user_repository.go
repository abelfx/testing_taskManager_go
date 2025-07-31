package mocks

import (
	"restfulapi/domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) (primitive.ObjectID, error) {
	args := m.Called(user)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *MockUserRepository) Authenticate(usernameOrEmail, password string) (*domain.User, error) {
	args := m.Called(usernameOrEmail, password)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) GetByID(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]*domain.User, error) {
	args := m.Called()
	users, _ := args.Get(0).([]*domain.User)
	return users, args.Error(1)
}

func (m *MockUserRepository) PromoteUser(userID primitive.ObjectID, newRole string) error {
	args := m.Called(userID, newRole)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteByID(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}
