package main

import "fmt"

func sumSubArray(arr []int, l int, r int) int {
	var sum int
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum
}

func MaxSequence(arr []int) int {

	var maxSqnc int = arr[0]
	var subArray int

	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			subArray = sumSubArray(arr, i, j)
			if maxSqnc < subArray {
				maxSqnc = subArray
			}
		}
	}
	return maxSqnc
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
