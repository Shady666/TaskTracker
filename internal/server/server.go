package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
	"taskTracker/config"
)

type Server struct {
	cfg       *config.Config
	app       *fiber.App
	mysqlConn *gorm.DB
}

func NewServer(cfg *config.Config, mysqlConn *gorm.DB) *Server {
	return &Server{
		cfg:       cfg,
		app:       fiber.New(fiber.Config{DisableStartupMessage: false}),
		mysqlConn: mysqlConn,
	}
}

// Run start server
func (s *Server) Run() {
	s.MapHandlers(s.cfg, s.mysqlConn)

	go func() {
		err := s.app.Listen(s.cfg.Port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("INTRO Shutting Down")

	if err := s.app.Shutdown(); err != nil {
		log.Printf("error occured on server shutting down: %s\n", err.Error())
	}

}
