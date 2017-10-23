package main

import (
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
	respond()
}

func request() {

}

func response() {
	res :=
		`<DOCTYPE html> 
			<html>
				<head>
				</head>
				<body>
					<strong>Hello World</strong>
				</body>
			</html>`
}
