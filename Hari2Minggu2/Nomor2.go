package main

import (
	"fmt"
)

func fibonacci(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 || number == 2 {
		return 1
	}
	if number > 2 {
		return fibonacci(number-2) + fibonacci(number-1)
	}
	return number
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
