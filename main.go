package main

import (
	"img_crawler/api"
	"net/http"
)

func main() {
	http.HandleFunc("/users/login", api.Login)
	http.HandleFunc("/pictures", api.GetPicture)
	http.HandleFunc("/", api.Index)
	http.ListenAndServe(":8080", nil)

}
