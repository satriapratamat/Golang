package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	prime := []int{}
	for i := 2; true; i++ {
		if primeNumber(i) {
			prime = append(prime, i)
		}
		if len(prime) == number {
			break
		}
	}
	return prime[number-1]

}

func primeNumber(number int) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
