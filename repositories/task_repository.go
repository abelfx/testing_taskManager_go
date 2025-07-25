package repositories

import (
	"context"
	"errors"
	"time"

	"restfulapi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
    Collection *mongo.Collection
}

func NewMongoTaskRepository(col *mongo.Collection) domain.TaskRepository {
    return &MongoTaskRepository{Collection: col}
}

func (r *MongoTaskRepository) Add(task *domain.Task) (primitive.ObjectID, error) {
    task.ID = primitive.NewObjectID()
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.Collection.InsertOne(ctx, task)
    return task.ID, err
}

func (r *MongoTaskRepository) GetAll() ([]*domain.Task, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var tasks []*domain.Task
    for cursor.Next(ctx) {
        var task domain.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &task)
    }
    return tasks, nil
}

func (r *MongoTaskRepository) GetByID(id primitive.ObjectID) (*domain.Task, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var task domain.Task
    err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
    if err != nil {
        return nil, err
    }
    return &task, nil
}

func (r *MongoTaskRepository) DeleteByID(id primitive.ObjectID) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("task not found")
    }
    return nil
}

func (r *MongoTaskRepository) UpdateByID(id primitive.ObjectID, task *domain.Task) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{
        "$set": bson.M{
            "title":       task.Title,
            "description": task.Description,
            "due_date":    task.DueDate,
            "status":      task.Status,
        },
    }

    result, err := r.Collection.UpdateByID(ctx, id, update)
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("task not found")
    }
    return nil
}
