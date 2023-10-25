package model

import (
	"fmt"

	"gorm.io/gorm"

	data "backend/model"
)

var db *gorm.DB

func CreateGroup(group *data.Group) error {
	// nilになってる
	fmt.Println(db)

	if err := db.Create(&group).Error; err != nil {
		return err
	}

	return nil
}
