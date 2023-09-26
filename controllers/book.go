package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RomainC75/postgres-test/models"
	"github.com/gin-gonic/gin"
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

func ListBooks(c *gin.Context) {
	var book []models.Book

	err := models.GetBooks(&book)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")

	if err := models.GetBookById(&book, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	}
	newBook := models.Book{
		Id: intId,
	}
	var oldBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.GetBookById(&oldBook, id); err != nil {
		log.Panicln(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := models.UpdateBook(&newBook, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, newBook)

}

func DeleteBook(c *gin.Context) {
	var foundBook models.Book
	id := c.Params.ByName("id")

	if err := models.GetBookById(&foundBook, id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := models.DeleteBook(id); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, nil)
}
