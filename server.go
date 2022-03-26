package main

import (
	"log"
	"net/http"
	"./handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/", handlers.SayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
