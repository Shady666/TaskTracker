package product

import "taskTracker/internal/models"

type TaskUC interface {
	CreateTask(task models.Task) (uint, error)
	UpdateTask(ID uint, description string) error
	DeleteTask(ID uint) error
	GetTaskByID(ID uint) (models.Task, error)
	FetchTasks() ([]models.Task, error)
}
