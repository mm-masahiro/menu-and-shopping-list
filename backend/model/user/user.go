package model

import (
	"gorm.io/gorm"

	data "backend/model"
)

var db *gorm.DB

func CreateUser(user *data.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
