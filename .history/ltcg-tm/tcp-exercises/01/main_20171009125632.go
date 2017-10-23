package main

import (
	"io"
	"log"
	"net/http"
)

// DefaultHandler prints "Hello world" to a blank query.
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

// DogHandler prints "woof woof" when dog is queried.
func DogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "woof woof!\n")
}

// DogHandler prints my name "Jonathan Wood" when "/me" queried.
func MeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Jonathan Wood\n")
}

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/dog/", DogHangler)
	http.HandleFunc("/me/", MeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
