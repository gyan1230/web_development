package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname, Lname string
	Age          int
}
type agent struct {
	person
	Ltk bool
}

func main() {
	http.HandleFunc("/mshl/", mshl)
	http.HandleFunc("/encode/", encd)
	http.ListenAndServe(":8080", nil)
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		Fname: "gyan",
		Lname: "chand",
		Age:   30,
	}

	json, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)

}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s1 := agent{
		person: person{
			Fname: "james",
			Lname: "bond",
			Age:   32,
		},
		Ltk: true,
	}
	err := json.NewEncoder(w).Encode(s1)
	if err != nil {
		log.Println(err)
	}

	// json, err := json.Marshal(s1)
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.Write(json)

}
