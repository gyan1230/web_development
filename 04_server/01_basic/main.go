package main

import (
	"fmt"
	"net/http"
)

type gyan struct{}

func (g gyan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	var g1 gyan
	http.ListenAndServe(":8080", g1)
}
