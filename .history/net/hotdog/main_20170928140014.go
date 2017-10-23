package main

import (
	"fmt"
)

type hotdog int

func (h hotdog) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {
	d hotdog
	http.ServeHTTP(":8080" d)
}
