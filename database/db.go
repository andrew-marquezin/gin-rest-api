package database

import (
	"gin-rest-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	connStr := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&models.Student{})
}
