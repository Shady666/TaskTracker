package repository

import (
	"gorm.io/gorm"

	"taskTracker/internal/models"
	"taskTracker/internal/product"
	"taskTracker/pkg/errorList"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) product.TaskRepo {
	return &TaskRepo{db: db}
}

func (t TaskRepo) CreateTask(task models.Task) (uint, error) {
	row := t.db.Create(&task)

	if err := row.Error; err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (t TaskRepo) UpdateTask(ID uint, description string) error {
	task := models.Task{ID: ID}

	result := t.db.First(&task)
	if result == nil {
		return errorList.ErrTaskNotExist
	}

	task.Description = description

	result = t.db.Save(&task)
	if result == nil {
		return errorList.ErrTaskNotUpdated
	}

	return nil
}

func (t TaskRepo) DeleteTask(ID uint) error {
	task := models.Task{ID: ID}

	result := t.db.Delete(&task)
	if result == nil {
		return errorList.ErrTaskNotExist
	}

	return nil
}

func (t TaskRepo) GetTaskByID(ID uint) (models.Task, error) {
	task := models.Task{ID: ID}

	result := t.db.First(&task)
	if result == nil {
		return models.Task{}, errorList.ErrTaskNotExist
	}

	return task, nil
}

func (t TaskRepo) FetchTasks() ([]models.Task, error) {
	var tasks []models.Task

	result := t.db.Find(&tasks)
	if result == nil {
		return []models.Task{}, errorList.ErrTasksNotExist
	}

	return tasks, nil
}
