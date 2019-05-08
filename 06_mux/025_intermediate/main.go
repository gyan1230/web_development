package main

import (
	"io"
	"net/http"
)

type gyan struct{}

func (g gyan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home":
		io.WriteString(w, "home..")
	case "/index":
		io.WriteString(w, "index.........")
	}

}

func main() {
	var g gyan
	http.Handle("/home", g)

	http.ListenAndServe(":8800", g)

}
