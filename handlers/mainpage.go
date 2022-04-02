package handlers

import (
	"fmt"
	"net/http"
	"../auth"
)

func MainMage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		statusToken, username := auth.VerifyToken(w, r)
		if statusToken != auth.TokenOK{
			if !auth.HandleToken(statusToken, username, w){
				return
			}
		}
		fmt.Fprintf(w, "Hello main page!")
	}else{
		http.Error(w, "", http.StatusBadRequest)
	}
}