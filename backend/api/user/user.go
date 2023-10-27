package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	data "backend/model"

	"gorm.io/gorm"
)

type CreateRequest struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	GroupId int    `json:"groupId"`
	Group   data.Group
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []data.User

	data.GetUsers((&users))

	responseBody, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var user data.User

	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}

	data.CreateUser((&user))

	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
