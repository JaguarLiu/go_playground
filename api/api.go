package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"img_crawler/auth"
	crawler "img_crawler/img_crawler"
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
		token, err := srv.CreateToken(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\"token\": \""+token+"\"}")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func GetPicture(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		imgs, err := crawler.Images()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
			return
		}
		json, err := json.Marshal(imgs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}
func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("views/index.html"))
	imgs, err := crawler.Images()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		return
	}
	tmpl.Execute(w, struct {
		Images []string
	}{
		imgs,
	})
}
