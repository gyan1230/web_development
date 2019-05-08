package main

import (
	"io"
	"net/http"
)

type gygn struct{}

func (g1 gygn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "In root")
	case "/home":
		io.WriteString(w, "In home")
	case "/about":
		io.WriteString(w, "About...")
	}

}

func main() {
	var g2 gygn
	http.ListenAndServe(":8080", g2)
}
