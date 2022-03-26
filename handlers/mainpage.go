package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"html/template"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"time"
)

func GenerateSalt(saltSize int)[]byte{
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil{
		return nil
	}
	return salt
}

func GenerateTocken(password []byte)string{
	salt := GenerateSalt(16)
	if salt == nil{
		return ""
	}
	var sha512 = sha512.New()
	var passwordWithSalt = make([]byte, 64)
	passwordWithSalt = append(passwordWithSalt, password...)
	passwordWithSalt = append(passwordWithSalt, salt...)
	sha512.Write(passwordWithSalt)
   	hashPassword := sha512.Sum(nil)
	var encodedHash = base64.URLEncoding.EncodeToString(hashPassword)
	return encodedHash
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./forms/mainpage.html")
		if err != nil{
			fmt.Println("Cant parse")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Server got some error!"))
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST"{
		r.ParseForm()
		password := r.FormValue("password")
		username := r.FormValue("username")
		if password == "" || username == ""{
			w.Write([]byte("You must enter name and password!"))
			return
		}
		token := GenerateTocken([]byte(password))
		if token == ""{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Server got some error!"))
			return
		}
		var dur time.Time = time.Now()
		dur.Add(time.Minute * 25)
		tokenCookie := &http.Cookie{Name: "SessionID", Value: token, HttpOnly: true}
		http.SetCookie(w, tokenCookie)

		fmt.Println("User:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
		

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

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}