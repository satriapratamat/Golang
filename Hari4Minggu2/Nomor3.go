package main

import (
	"fmt"
	"math"
)

var res = make(map[int]int)

func jump(a, b int) int {
	frogJump := math.Abs((float64(a - b)))
	fmt.Println("frogJump line12 adalah = ", frogJump)
	return int(frogJump)
}

func Frog(jumps []int) int {

	nJump := len(jumps) - 1
	fmt.Println("nJump line19 adalah = ", nJump)
	res[0] = 0
	res[1] = jump(jumps[0], jumps[1])
	for i := 2; i <= nJump; i++ {
		min2 := float64(res[i-2] + jump(jumps[i], jumps[i-2]))
		// fmt.Println("min2 adalah = ", min2)
		min1 := float64(res[i-1] + jump(jumps[i], jumps[i-1]))
		// fmt.Println("min1 adalah = ", min1)
		res[i] = int(math.Min(min2, min1))
		fmt.Println("res i adalah = ", res[i])
	}

	return res[nJump]

}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
