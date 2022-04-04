package handlers
import (
	"fmt"
	"net/http"
	"encoding/json"
	"../auth"
)

func RegisterGet(w http.ResponseWriter, r *http.Request){
	_, err := r.Cookie("token")
	if err != nil {
		fmt.Fprintf(w, "Hello register page with now cookies!")
	}
	fmt.Fprintf(w, "Hello register page!")
}

func RegisterPost(w http.ResponseWriter, r *http.Request){
	var creds auth.AuthCreds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil{
		http.Error(w, "You must enter name and password using JSON!", http.StatusBadRequest)
		return
	}
	if creds.Password == "" || creds.User == ""{
		http.Error(w, "You must enter name and password", http.StatusBadRequest)
		return;
	}
	if !auth.RegisterUser(creds.User, creds.Password){
		http.Error(w, "Error, perhaps username already registered", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Signed up user %v", creds.User)
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