package main

import (
	"net/http"

	"gorm.io/gorm"

	group "backend/api/group"
	user "backend/api/user"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primarykey"`
	Name     string
	Email    string
	IsActive bool
}

func main() {
	http.HandleFunc("/users", user.GetUsers)
	http.HandleFunc("/user/create", user.Create)
	http.HandleFunc("/group/create", group.Create)
	http.ListenAndServe(":8080", nil)
}
