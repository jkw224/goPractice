package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Response) {
	fmt.Fprintln(w, "Any code you want here")
}

func main() {
	var d hotdog
	http.ServeHTTP(":8080", d)
}
