package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Response)

func main() {
	d hotdog
	http.ServeHTTP(":8080", d)
}
