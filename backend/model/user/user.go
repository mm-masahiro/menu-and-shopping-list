package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool
}

var db *gorm.DB

func CreateUser(user *User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func init() {
	var err error

	godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"), db_user, db_password, db_name)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}
