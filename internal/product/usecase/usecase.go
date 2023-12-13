package usecase

import (
	"taskTracker/internal/models"
	"taskTracker/internal/product"
)

type TaskUC struct {
	repo product.TaskRepo
}

func NewTaskUC(repo product.TaskRepo) product.TaskUC {
	return &TaskUC{repo: repo}
}

func (t TaskUC) CreateTask(task models.Task) (uint, error) {
	return t.repo.CreateTask(task)
}

func (t TaskUC) UpdateTask(ID uint, description string) error {
	return t.repo.UpdateTask(ID, description)
}

func (t TaskUC) DeleteTask(ID uint) error {
	return t.repo.DeleteTask(ID)
}

func (t TaskUC) GetTaskByID(ID uint) (models.Task, error) {
	return t.repo.GetTaskByID(ID)
}

func (t TaskUC) FetchTasks() ([]models.Task, error) {
	return t.repo.FetchTasks()
}
