package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool
}

type Group struct {
	gorm.Model
	ParentUserId int `json:"parentUserId"`
}
