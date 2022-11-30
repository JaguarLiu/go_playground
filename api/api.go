package api

import (
	"encoding/json"
	"fmt"
	"img_crawler/auth"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		account := r.PostForm.Get("account")
		password := r.PostForm.Get("password")
		user := auth.User{Account: account,
			Password: password,
		}
		srv := auth.Auth{}
		err := srv.Login(user)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
func GetPicture(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}
