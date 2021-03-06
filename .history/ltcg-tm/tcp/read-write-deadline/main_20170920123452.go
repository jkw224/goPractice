package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Split()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %v", ln)
	}
	defer conn.Close()

	// now we get here
	// the connection times out
	// that breaks us out of the scanner loop
	fmt.Println("*** CODE GOT HERE ***")
}
