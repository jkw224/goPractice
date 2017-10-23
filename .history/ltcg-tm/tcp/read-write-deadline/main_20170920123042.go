package main

import (
	"time"
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
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	scanner := bufio.NewScanner(conn)
	bufio.
}
