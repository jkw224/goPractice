package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// returns listener & error
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Error reading request on port :8080...", err)
	}
	defer li.Close()

	for {
		// returns Conn & error -- returns next connection to listener
		conn, err := li.Accept()
		if err != nil {
			log.Println("li.Accept() returns no connection from port :8080...", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// Never get here
	// b/c have an open connection
	// how does above reader know when it's done?
	fmt.Println("Code got here.")
}
