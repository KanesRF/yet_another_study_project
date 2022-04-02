package handlers

import (
	"net/http"
	"../auth"
	"../db"
)

func Logout(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		statusToken, username := auth.VerifyToken(w, r)
		if statusToken != auth.TokenOK{
			if !auth.HandleToken(statusToken, username, w){
				return
			}
		}
		rows, err := db.DbConn.Query("UPDATE public.users set signed_in = $1", false)
		if err!= nil{
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
	}else{
		http.Error(w, "", http.StatusBadRequest)
	}
}