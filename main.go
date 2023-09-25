package main

import (
	"log"
	"net"
	"net/http"

	"github.com/RomainC75/postgres-test/db"
	"github.com/RomainC75/postgres-test/models"
	"github.com/RomainC75/postgres-test/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.POST("/book", routes.AddBook)
	mux.GET("/book", routes.FindBook)

	db := db.Init()
	db.AutoMigrate(&models.Book{})

	// newBook := models.Book{

	// 	Title:  "mlsdkjf",
	// 	Author: "lkfnons",
	// 	Desc:   "jjjjjjj",
	// }
	// db.Create(&newBook)

	log.Println("Starting server....")
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, mux)
}
