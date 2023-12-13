package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(group fiber.Router, handlers TaskHandlers) {
	group.Post("/create", handlers.CreateTask)
	group.Post("/update", handlers.UpdateTask)
	group.Get("/get", handlers.GetTaskByTD)
	group.Get("/fetch", handlers.FetchTask)
	group.Delete("/delete", handlers.DeleteTask)
}
