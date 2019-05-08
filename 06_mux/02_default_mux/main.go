package main

import (
	"io"
	"net/http"
)

type gyan struct{}

func (g gyan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog/":
		io.WriteString(w, "dog ... woof")
	case "/cat":
		io.WriteString(w, "Cat ... Meow")
	}

}

func main() {
	var g gyan
	myMux := http.NewServeMux()
	myMux.Handle("/dog/", g)
	myMux.Handle("/cat", g)
	http.ListenAndServe(":8080", myMux)

}
