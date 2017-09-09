package main

import (
	"fmt"
)

func main() {

	var n []int

	for i := 0; i <= 10; i++ {
		n = append(n, i)
	}

	for _, num := range n {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
