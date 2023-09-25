package models

import (
	"errors"
	"fmt"

	"github.com/RomainC75/postgres-test/db"
)

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func CreateBook(b *Book) (*Book, error) {

	fmt.Print("received : ", *b, "\n")
	// book := &Book{}
	res := db.DB.Create(&b)

	if res.RowsAffected == 0 {
		return nil, errors.New("error saving todo")
	}
	fmt.Print("-> affected", res)
	// return &pb.Todo{
	// 	Id:     todo.Id,
	// 	Title:  todo.Name,
	// 	Author: todo.Description,
	// 	Desc:   false,
	// }, nil
	return nil, nil
}

func GetBooks(b *[]Book) error {
	if err := db.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
