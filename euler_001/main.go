package main

import (
	"fmt"
)

func main() {
	var target int = 1000

	sum := 0
	for i := 1; i < target; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
