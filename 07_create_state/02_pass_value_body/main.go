package main

import (
	"fmt"
	"net/http"
)

type person struct {
	first, last string
	license     bool
	sal         float32
	age         int
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	b := r.FormValue("license") == "on"
	s := r.FormValue("sal")
	a := r.FormValue("age")

	//fmt.Println("First: ", f, "\nLast: ", l, "\nLicense: ", b, "\nSalary: ", s, "\nAge: ", a)
	fmt.Fprintln(w, "value :", f, l, b, s, a)
}
