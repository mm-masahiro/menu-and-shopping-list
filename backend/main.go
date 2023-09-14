package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(("%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"), db_user, db_password, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	http.HandleFunc("/user/create", user.Create)
	http.ListenAndServe(":8080", nil)
}
