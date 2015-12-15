package main

import (
	"fmt"
	"math"
)

func is_prime(input int) bool {
	var test int = int(math.Sqrt(float64(input)))

	for test > 1 {
		if input%test == 0 {
			return false
		}

		test -= 1
	}

	return true
}

func main() {
	var input float64 = 600851475143

	var test int = int(math.Sqrt(input))

	for int(input)%test != 0 || !is_prime(test) {
		test -= 1
	}

	fmt.Println(test)
}
