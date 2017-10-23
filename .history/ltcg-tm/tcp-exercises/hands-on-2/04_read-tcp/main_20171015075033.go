package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer c.Close()
	}
	// we never get her
	// we have an open stream connection
	// how does the above read know when it's done
	fmt.Println("Code got here.")
}
