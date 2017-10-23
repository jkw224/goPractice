package main

import (
	"fmt"
)

func main() {
	// send multiple #s into multiple channels
	c := gen()
	//
	f := factorial(c)
	for n := range f {
		fmt.Println(n)
	}
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
