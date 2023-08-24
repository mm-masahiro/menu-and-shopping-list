package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "Hello, world!!!")
}
