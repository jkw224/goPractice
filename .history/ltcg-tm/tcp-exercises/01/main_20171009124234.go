package main

import (
	"io"
	"log"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

func DogHangler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "woof woof!\n")
}

func MeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Jonathan Wood\n")
}

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/dog/")
	http.HandleFunc("/me/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
