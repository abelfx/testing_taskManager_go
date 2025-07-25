package usecases

import (
	"errors"

	"restfulapi/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
    repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
    return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) CreateTask(task *domain.Task) (primitive.ObjectID, error) {
    return u.repo.Add(task)
}

func (u *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
    return u.repo.GetAll()
}

func (u *TaskUsecase) GetTaskByID(id string) (*domain.Task, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid id")
    }
    return u.repo.GetByID(objID)
}

func (u *TaskUsecase) DeleteTaskByID(id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid id")
    }
    return u.repo.DeleteByID(objID)
}

func (u *TaskUsecase) UpdateTask(id string, task *domain.Task) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid id")
    }
    return u.repo.UpdateByID(objID, task)
}
