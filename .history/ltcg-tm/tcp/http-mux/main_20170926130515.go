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
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		handle(conn)
	}
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func handle(conn net.Conn) {

}
