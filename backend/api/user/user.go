package api

import (
	"fmt"
	"net/http"
)

type CreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println(string(body))
}
