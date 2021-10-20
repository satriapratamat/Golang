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

	fmt.Println(primeNumber(11)) // true

	fmt.Println(primeNumber(13)) // true

	fmt.Println(primeNumber(17)) // true

	fmt.Println(primeNumber(20)) // false

	fmt.Println(primeNumber(35)) // false

}
