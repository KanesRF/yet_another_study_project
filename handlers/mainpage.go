package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"../user"
	"time"
)




func LoginPOST(w http.ResponseWriter, r *http.Request){
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
	fmt.Println(creds)
	if creds.Password == "" || creds.User == ""{
		http.Error(w, "You must enter name and password", http.StatusBadRequest)
		return;
	}
	if !user.AuthByPassword(creds.Password, creds.User){
		http.Error(w, "Wrong username or password", http.StatusUnauthorized)
		return
	}
	accessTocken, err := user.GenerateJwtTocken(time.Now().Add(5 * time.Minute), creds.User)
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
	fmt.Println(accessTocken)

}

func LoginGET(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			fmt.Fprintf(w, "Hello login page with now cookies!")
		}
		fmt.Fprintf(w, "Hello login page! %v", tokenCookie.Value)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET":
		LoginGET(w, r)
	case "POST":
		LoginPOST(w, r)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./forms/mainpage.html")
		if err != nil{
			fmt.Println("Cant parse")
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("User:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
	}
}

func MainMage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			fmt.Println("No cookie")
			fmt.Fprintf(w, "Hello main page with now cookies!")
			return;
		}
		fmt.Println(tokenCookie)
		fmt.Fprintf(w, "Hello main page!")
	}
}