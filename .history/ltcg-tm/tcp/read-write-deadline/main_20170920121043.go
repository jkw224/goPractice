package main

import (
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {

	}
	defer li.Close()

}
