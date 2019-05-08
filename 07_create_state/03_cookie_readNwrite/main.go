package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:  "my-cookie",
		Value: "asdf123",
	}
	http.SetCookie(w, &c)
}

func read(w http.ResponseWriter, r *http.Request) {
	get, err := r.Cookie("my-cookie")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, "cookie is: ", get)
}
