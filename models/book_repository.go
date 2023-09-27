package models

import (
	"time"

	"github.com/RomainC75/postgres-test/db"
)

func CreateBook(b *Book) (*Book, error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	if err := db.DB.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetBooks(b *[]Book) error {
	if err := db.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetBookById(b *Book, key string) error {
	if err := db.DB.First(&b, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(b *Book, key string) error {
	// need the complete element ?
	b.UpdatedAt = time.Now()
	db.DB.Save(&b)
	return nil
}

func DeleteBook(key string) error {
	db.DB.Delete(&Book{}, key)
	return nil
}
