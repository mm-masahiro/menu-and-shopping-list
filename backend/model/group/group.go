package model

import (
	"gorm.io/gorm"

	data "backend/model"
)

var db *gorm.DB

func CreateGroup(group *data.Group) error {
	if err := db.Create(&group).Error; err != nil {
		return err
	}

	return nil
}
