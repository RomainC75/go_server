package models

import (
	"fmt"
	"time"

	"github.com/RomainC75/postgres-test/db"
)

func CreateUser(u *User) (*User, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if err := db.DB.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(u *[]User) error {
	if err := db.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(u *User, username string) error {
	fmt.Print("-> \ninside getUserByUsername ", *u)
	fmt.Print("\n-> \ninside getUserByUsername ", username)
	if err := db.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(u *User, key string) error {
	// need the complete element ?
	u.UpdatedAt = time.Now()
	db.DB.Save(&u)
	return nil
}
