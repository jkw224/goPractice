package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index returns index when no routes provided
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world")
}

// Me returns my name for /me/
func Me(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Jonathan Wood")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/me/", Me)
	log.Fatal(http.ListenAndServe(":8080", router))
}
