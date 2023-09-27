package models

import (
	"time"
)

type Book struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:false"`
}

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:false"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
