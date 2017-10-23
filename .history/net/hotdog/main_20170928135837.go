package main

import (
	"fmt", 
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {

}
