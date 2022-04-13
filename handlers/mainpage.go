package handlers

import (
	"fmt"
	"ga_server/auth"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	statusToken, username := auth.VerifyToken(w, r)
	if statusToken != auth.TokenOK {
		if !auth.HandleToken(statusToken, username, w) {
			return
		}
	}
	fmt.Fprintf(w, "Hello main page!")
}
