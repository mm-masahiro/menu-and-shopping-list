package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	data "backend/model"
	model "backend/model/user"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateRequest struct {
	Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var user data.User

	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}

	model.CreateUser((&user))

	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
