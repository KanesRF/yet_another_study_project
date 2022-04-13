package handlers

import (
	"encoding/json"
	"fmt"
	"ga_server/auth"
	"ga_server/db"
	"net/http"
	"time"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	defer func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	}(w, r)
	var creds auth.AuthCreds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "You must enter name and password using JSON!", http.StatusBadRequest)
		return
	}
	if creds.Password == "" || creds.User == "" {
		http.Error(w, "You must enter name and password", http.StatusBadRequest)
		return
	}
	if !auth.AuthByPassword(creds.Password, creds.User) {
		http.Error(w, "Wrong username or password", http.StatusUnauthorized)
		return
	}
	accessTocken, err := auth.GenerateJwtTocken(time.Now().Add(auth.TokenLifeTime), creds.User)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	rows, err := db.DbConn.Query("UPDATE public.users set signed_in = $1", true)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    accessTocken,
		HttpOnly: true,
	})
	fmt.Fprintf(w, "Successfully signed in")
}

func LoginGet(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err != nil {
		fmt.Fprintf(w, "Hello login page with cookies!")
	}
	fmt.Fprintf(w, "Hello login page!")

}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		LoginGet(w, r)
	case "POST":
		LoginPost(w, r)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
