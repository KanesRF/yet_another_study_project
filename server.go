package main

import (
	"log"
	"net/http"
	"./handlers"
	"./db"
)

func main() {
	db.InitDB()
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/", handlers.MainMage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	db.CloseDB()
}
