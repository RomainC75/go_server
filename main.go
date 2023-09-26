package main

import (
	"log"

	"github.com/RomainC75/postgres-test/db"
	"github.com/RomainC75/postgres-test/models"
	"github.com/RomainC75/postgres-test/routes"
	"github.com/gin-gonic/gin"
)

var err error

func main() {
	// mux := httprouter.New()

	db := db.Init()
	db.AutoMigrate(&models.Book{})

	// newBook := models.Book{

	// 	Title:  "mlsdkjf",
	// 	Author: "lkfnons",
	// 	Desc:   "jjjjjjj",
	// }
	// db.Create(&newBook)

	log.Println("Starting server....")
	// listener, err := net.Listen("tcp", ":8081"

	// http.Serve(listener, mux)

	r := routes.SetupRouter()
	r.Use(gin.Logger())

	err = r.Run()
}
