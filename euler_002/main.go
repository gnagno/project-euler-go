package main

import (
	"fmt"
)

func main() {
	var target int = 4000000

	var prev int = 1
	var current int = 2
	var sum int = 0

	for current < target {
		if current%2 == 0 {
			sum += current
		}

		prev, current = current, prev+current
	}

	fmt.Println(sum)
}
