package main

import (
	"net"
)

func main() {
	ln, err := net.Listen(tcp, ":8080")
}
