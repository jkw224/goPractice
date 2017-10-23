package main

import (
	"fmt"
)

type hotdog int

func ( hotdog) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {
	h 
}
