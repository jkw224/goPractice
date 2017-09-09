package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type file []string

func main() {
	// https://gobyexample.com/command-line-arguments
	// arguments passed in via the command line
	// first value in slice = path to the program
	//
	// allArgs := os.Args

	// for i, arg := range allArgs {
	// 	fmt.Printf("arg %v: %v type:%T\n", i, arg, arg)
	// }

	var firstArg file = os.Args[1:]
	f, err := os.Open(firstArg.toString())
	check(err)

	bs100 := make([]byte, 100)
	n, err := f.Read(bs100) // n = # of bytes read
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(bs100))
}

func (ss file) toString() string {
	s := strings.Join(ss, "")
	return s
	// return trimNonAlpha(s)
}

func trimNonAlpha(s string) string {
	// returns only alphanumeric + "." characters
	// e.g. filename.txt
	reg, err := regexp.Compile("[^a-zA-Z0-9.]+")
	fatalCheck(err)
	processedString := reg.ReplaceAllString(s, "")

	fmt.Printf("A string of %s becomes %s \n", s, processedString)
	return processedString
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fatalCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
