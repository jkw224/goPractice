package main

import (
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(ResponseWriter)

func main() {

}