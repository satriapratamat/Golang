package main

import "fmt"

func fibo(n int) int {

	// your code here
	// jika n = 0, fibo(n) = 0
	// jika n = 1, fibo(n) = 1
	// pada kasus ini, pada n = 0 dan n = 1, n = fibo(n)

	if n <= 1 {
		return n
	}

	n0 := 0 // nilai fibo(n) jika n = 0
	n1 := 1 // nilai fibo(n) jika n = 1

	//jika n = 2, fibo(n) = fibo(n-2) + fibo(n-1)
	var nX int

	for i := 2; i <= n; i++ {
		nX = n0 + n1
		n0 = n1
		n1 = nX
	}
	return nX

}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 13

	fmt.Println(fibo(10)) // 55

}
