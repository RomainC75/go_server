package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	// DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	url := "postgres://name:pass@localhost:5432/mydb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	log.Println("db connection successful")

	DB = db
	return db
}
