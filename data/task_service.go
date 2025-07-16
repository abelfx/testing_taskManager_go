package data

import "restfulapi/models"

// adds a task
func Add(task *models.Task) bool {
	for _,existingTask := range models.Tasks {
		if task.ID == existingTask.ID {
			return false
		}
	}
	models.Tasks = append(models.Tasks, task)
	return true
}

// gets all avaliable tasks
func GetAll() []*models.Task{
	return models.Tasks
}

// gets a specific task using its id
func Get(id string) (*models.Task, bool) {
	for _, task := range models.Tasks {
		if task.ID == id {
			return task, true
		}
	}
	return nil, false
}


// delete a task using its id
func Delete(id string) bool{
	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return true
		}
	}
	return false

}

// update a task using its id
func Update(id string, updatedTask *models.Task) bool {
	for i, task := range models.Tasks {
		if task.ID == id {
			// Preserve the ID from the URL and update other fields
			models.Tasks[i].Title = updatedTask.Title
			models.Tasks[i].Description = updatedTask.Description
			models.Tasks[i].DueDate = updatedTask.DueDate
			models.Tasks[i].Status = updatedTask.Status
			return true
		}
	}
	return false
}