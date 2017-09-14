package main

import (
	"fmt"
)

func main() {
	// Saved the Rune (int32) a into foo
	var foo = 'a'

	fmt.Println(foo)
	fmt.Printf("\n%T\n\n", foo)

	for i := 32; i < 100; i++ {
		fmt.Printf("%v - %v - %v - %b\n", i, string(i), []byte(string(i)), i)
	}
}
