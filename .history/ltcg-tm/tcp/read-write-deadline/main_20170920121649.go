package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	conn, err := li.Accept()
	if err != nil {
		fmt.Println(err)
	}

	handle(conn)
}

func handle(conn net.Conn) {
	bufio.NewScanner(conn)
}
