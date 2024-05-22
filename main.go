package main

import (
	"net/http"
	"vikash/handler"
)

func main() {
	http.HandleFunc("/user/create", handler.CreateUser)
	http.ListenAndServe(":8080", nil)
}
