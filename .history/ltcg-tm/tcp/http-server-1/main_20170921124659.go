package main

import (
	"strings"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err.Error)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for bufio.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i ==0 {
			// request line
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
			if ln == "" {
				// headers are done
				break
			}
			i++
		}
	}
}

func response(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>Hello world!</strong>
			</body>
		</html>`
	
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}