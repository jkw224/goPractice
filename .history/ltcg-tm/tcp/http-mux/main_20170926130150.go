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
		conn, err := ln.Accept() {

		}
		conn.Close()
	}	
}