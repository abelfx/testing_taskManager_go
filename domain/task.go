package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task entity
type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    DueDate     time.Time          `bson:"due_date" json:"due_date"`
    Status      string             `bson:"status" json:"status"`
}

// TaskRepository interface abstracts DB operations
type TaskRepository interface {
    Add(task *Task) (primitive.ObjectID, error)
    GetAll() ([]*Task, error)
    GetByID(id primitive.ObjectID) (*Task, error)
    DeleteByID(id primitive.ObjectID) error
    UpdateByID(id primitive.ObjectID, task *Task) error
}
