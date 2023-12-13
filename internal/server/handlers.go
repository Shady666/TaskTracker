package server

import (
	"gorm.io/gorm"
	"taskTracker/config"
	handlers "taskTracker/internal/product/delivery/http"
	"taskTracker/internal/product/repository"
	"taskTracker/internal/product/usecase"
)

// MapHandlers init all fiber handleFuncs
func (s *Server) MapHandlers(cfg *config.Config, mysqlConn *gorm.DB) {
	taskRepos := repository.NewTaskRepo(mysqlConn)
	taskUC := usecase.NewTaskUC(taskRepos)
	taskHandlers := handlers.NewTaskHandler(s.cfg, taskUC)

	taskGroup := s.app.Group("tasks")

	handlers.MapRoutes(taskGroup, *taskHandlers)
}
