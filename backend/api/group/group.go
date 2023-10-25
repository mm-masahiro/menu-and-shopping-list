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
	ParentUserId int `json:"parentUserId"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var group data.Group

	if err := json.Unmarshal(reqBody, &group); err != nil {
		log.Fatal(err)
	}

	data.CreateGroup((&group))

	responseBody, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
