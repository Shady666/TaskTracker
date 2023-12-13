package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"taskTracker/config"
	"taskTracker/internal/models"
	"taskTracker/internal/product"
)

type TaskHandlers struct {
	cfg *config.Config
	uc  product.TaskUC
}

func NewTaskHandler(cfg *config.Config, uc product.TaskUC) *TaskHandlers {
	return &TaskHandlers{cfg: cfg, uc: uc}
}

func (t TaskHandlers) CreateTask(c *fiber.Ctx) error {
	var task models.Task

	err := json.Unmarshal(c.Body(), &task)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	id, err := t.uc.CreateTask(task)
	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(models.Response{Message: err.Error()})
	}

	return c.JSON(models.Response{ID: id})
}

func (t TaskHandlers) UpdateTask(c *fiber.Ctx) error {
	var task models.Task

	err := json.Unmarshal(c.Body(), &task)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = t.uc.UpdateTask(task.ID, task.Description)
	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(models.Response{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t TaskHandlers) GetTaskByTD(c *fiber.Ctx) error {
	result, err := t.uc.GetTaskByID(uint(c.QueryInt("id")))
	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(models.Response{Message: err.Error()})
	}

	return c.JSON(result)
}

func (t TaskHandlers) DeleteTask(c *fiber.Ctx) error {
	err := t.uc.DeleteTask(uint(c.QueryInt("id")))
	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(models.Response{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t TaskHandlers) FetchTask(c *fiber.Ctx) error {
	result, err := t.uc.FetchTasks()
	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(models.Response{Message: err.Error()})
	}

	return c.JSON(result)
}
