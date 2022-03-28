package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../user"
	"time"
)

func LoginPost(w http.ResponseWriter, r *http.Request){
	defer func(w http.ResponseWriter, r *http.Request) {
        if err := recover(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
        }
    }(w, r)
	var creds user.AuthCreds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil{
		http.Error(w, "You must enter name and password using JSON!", http.StatusBadRequest)
		return
	}
	if creds.Password == "" || creds.User == ""{
		http.Error(w, "You must enter name and password", http.StatusBadRequest)
		return;
	}
	if !user.AuthByPassword(creds.Password, creds.User){
		http.Error(w, "Wrong username or password", http.StatusUnauthorized)
		return
	}
	accessTocken, err := user.GenerateJwtTocken(time.Now().Add(user.TokenLifeTime), creds.User)
	if err != nil{
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   accessTocken,
		HttpOnly : true,
	})
	fmt.Fprintf(w,"Successfully signed in")
}

func LoginGet(w http.ResponseWriter, r *http.Request){
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		fmt.Fprintf(w, "Hello login page with now cookies!")
	}
	fmt.Fprintf(w, "Hello login page! %v", tokenCookie.Value)
	
}

func RegisterGet(w http.ResponseWriter, r *http.Request){
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		fmt.Fprintf(w, "Hello register page with now cookies!")
	}
	fmt.Fprintf(w, "Hello register page! %v", tokenCookie.Value)
}

func RegisterPost(w http.ResponseWriter, r *http.Request){
	var creds user.AuthCreds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil{
		http.Error(w, "You must enter name and password using JSON!", http.StatusBadRequest)
		return
	}
	if creds.Password == "" || creds.User == ""{
		http.Error(w, "You must enter name and password", http.StatusBadRequest)
		return;
	}
	if !user.RegisterUser(creds.User, creds.Password){
		http.Error(w, "Username already registered", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Signed up user %v", creds.User)
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		LoginGet(w, r)
	case "POST":
		LoginPost(w, r)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		RegisterGet(w, r)
	case "POST":
		RegisterPost(w, r)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func MainMage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	statusToken, username := user.VerifyToken(w, r)
	if statusToken != user.TokenOK{
		user.HandleToken(statusToken, username, w)
	}
	fmt.Fprintf(w, "Hello main page!")
	}else{
		http.Error(w, "", http.StatusBadRequest)
	}
}