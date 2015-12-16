package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(input int) bool {
	str := strconv.Itoa(input)
	start := 0
	end := len(str) - 1

	for start < end {
		if str[start] != str[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	max := 0

	for x := 999; x > 900; x-- {
		for y := 999; y > 900; y-- {
			test := x * y
			if test > max && isPalindrome(test) {
				max = test
			}
		}
	}

	fmt.Println(max)
}
