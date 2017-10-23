package main

import (
	"fmt"
	"net/http"
)

type hotdog int



func main() {
	d hotdog
	http.ServeHTTP(":8080", d)
}
