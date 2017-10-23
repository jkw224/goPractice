package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 8080 on all available unicast and
	// anycast IP addresses of the local system.
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		// Wait for a connection.
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			io.WriteString()
			c.Close(w, "I see you")
		}(conn)
	}
}
