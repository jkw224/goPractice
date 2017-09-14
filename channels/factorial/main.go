package main

import (
	"fmt"
)

func main() {
	for n := range factorial(4) {
		fmt.Println(n)
	}
}

func factorial(num int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for num > 0 {
			total *= num
			num--
		}
		out <- total
		close(out)
	}()
	return out
}
