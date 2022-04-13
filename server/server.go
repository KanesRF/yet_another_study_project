package main

import (
	"ga_server/db"
	"ga_server/handlers"
	"log"
	"net/http"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

func main() {
	db.InitDB()
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/", handlers.MainPage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	db.CloseDB()
}
