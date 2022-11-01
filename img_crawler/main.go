package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("https://ferrari-cdn.thron.com/api/xcontents/resources/delivery/getThumbnail/ferrari/0x640/53446961-5c54-49d4-9e2a-92bdfc90497b.jpg?v=102")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create("../asset/dream_car.jpg")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		panic(err)
	}
}
