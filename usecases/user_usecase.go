package usecases

import (
	"errors"

	"restfulapi/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(user *domain.User) (primitive.ObjectID, error) {
	return u.repo.Create(user)
}

func (u *UserUsecase) Authenticate(usernameOrEmail, password string) (*domain.User, error) {
	return u.repo.Authenticate(usernameOrEmail, password)
}

func (u *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) PromoteUser(userID primitive.ObjectID, newRole string) error {
	if newRole == "" {
		return errors.New("new role must be provided")
	}
	return u.repo.PromoteUser(userID, newRole)
}

func (u *UserUsecase) DeleteUserByID(id primitive.ObjectID) error {
	return u.repo.DeleteByID(id)
}
