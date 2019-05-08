package main

import (
	"io"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/about":
		io.WriteString(w, "In about")
	case "/contact/":
		io.WriteString(w, "Contact us ....")
	}
}

func main() {
	http.HandleFunc("/about", myHandler)
	http.HandleFunc("/contact/", myHandler)

	http.ListenAndServe(":8080", nil)
}
