package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {

	for i := 2; i <= int(math.Floor(float64(number)/2)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

func main() {

	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true

}
