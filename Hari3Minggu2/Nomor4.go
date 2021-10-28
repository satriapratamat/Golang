package main

import "fmt"

func BinarySearch(array []int, x int) {

	low, high := 0, len(array)-1

	if x > array[len(array)-1] || x < array[0] {
		fmt.Println(-1)
		return
	}
	if array[low] == x {
		fmt.Println(low)
		return
	}
	if array[high] == x {
		fmt.Println(high)
		return
	} else {
		for low+1 != high {
			if array[(low+high)/2] < x {
				low = (low + high) / 2
			}
			if array[(low+high)/2] == x {
				fmt.Println((low + high) / 2)
				return
			} else {
				high = (low + high) / 2
			}
		}
	}

	fmt.Println(-1)
	return
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
