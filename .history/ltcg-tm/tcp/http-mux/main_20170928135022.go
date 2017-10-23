package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new go routine.
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Handles request
	request(conn)

	// Handles response
	respond(conn)
}

func request(conn net.Conn) {

}

func respond() {
	body :=
		`<!DOCTYPE html> 
		<html lang="en">
			<head>
				<meta charet="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>Hello World</strong>
			</body>
		</html>`

	fmt.Fprint()

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
}
