package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RomainC75/postgres-test/models"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if _, err := models.CreateBook(&book); err != nil {
		log.Println(err.Error())
	}
	fmt.Print(book)

	c.JSON(http.StatusOK, book)

}

func ListBooks(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var book []models.Book

	err := models.GetBooks(&book)
	if err != nil {
		fmt.Printf("abort")
	}

	// fmt.Fprintf(w, "%+v", book)

	by, err := json.Marshal(book)
	if err != nil {
		fmt.Print("marshall error : ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(by)
}

func GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var book models.Book
	id := p.ByName("id")
	fmt.Print("-> ", id)
	if err := models.GetBookById(&book, id); err != nil {
		fmt.Print("get by id err: ", err)
	}

	by, err := json.Marshal(book)
	if err != nil {
		fmt.Print("marshall error : ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(by)

}
