package pgConnector

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taskTracker/internal/models"
)

func Connect() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
		return nil
	}

	dsn := "host=db-postgres user=postgres password=password123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())

		return nil
	}

	log.Println("connection status ok")

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf(err.Error())

		return nil
	}

	return db
}
