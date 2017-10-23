package main

import (
	"log"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err := nil {
			log.Panic(err)
		}
		hangle(conn)
	}	
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func handle() {

}