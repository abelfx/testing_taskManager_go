package data

import (
	"context"
	"restfulapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// adds a task
func Add(task *models.Task) (primitive.ObjectID, error) {
	task.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := TaskCollection.InsertOne(ctx, task)
	return task.ID, err

}

// gets all avaliable tasks
func GetAll() ([]*models.Task, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err:= cursor.Decode(&task); err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}
	return tasks, nil

}

// gets a specific task using its id
func Get(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task models.Task
	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil

}


// delete a task using its id
func Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err


}

// update a task using its id
func Update(id string, updatedTask *models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M {
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	_, err = TaskCollection.UpdateByID(ctx, objID, update)
	return err
}