package main

import (
	"net/http"

	"gorm.io/gorm"

	user "backend/api/user"
	models "backend/model"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primarykey"`
	Name     string
	Email    string
	IsActive bool
}

func main() {
	models.InitDb()

	http.HandleFunc("/user/create", user.Create)
	http.ListenAndServe(":8080", nil)
}
