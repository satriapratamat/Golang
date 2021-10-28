package main

import "fmt"

func SimpleEquations(a, b, c int) (x, y, z int) {
	eq := (a + b + c) / 3
	for i := 1; i <= eq; i++ {
		for j := 1; j <= eq; j++ {
			for k := 1; k <= eq; k++ {
				if i+j+k == a {
					if i*j*k == b {
						if i*i+j*j+k*k == c {
							fmt.Println(i, j, k)
							return
						}
					}
				}
			}
		}
	}

	fmt.Println("no solution")
	return
}

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

}
