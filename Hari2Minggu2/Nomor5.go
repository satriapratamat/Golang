package main

import (
	"fmt"
	"strconv"
)

func FindMinAndMax(arr []int) string {

	maxIndex, max := 0, 0
	for index, val := range arr {
		if max > val {
			continue
		}
		if max < val {
			maxIndex = index
			max = val
		}
	}
	minIndex, min := 0, 0
	for index, val := range arr {
		if max < val {
			continue
		}
		if max > val {
			minIndex = index
			min = val
		}
	}
	minNum := strconv.Itoa(min)
	maxNum := strconv.Itoa(max)
	minIdxNum := strconv.Itoa(minIndex)
	maxIdxNum := strconv.Itoa(maxIndex)
	result := "min: " + minNum + " index: " + minIdxNum + " max: " + maxNum + " index: " + maxIdxNum

	return result
}

func main() {

	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	// min: -2 index: 3 max: 8 index: 5
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	// min: -5 index: 1 max: 22 index: 3
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	// min: -21 index: 4 max: 9 index: 2
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	// min: -1 index: 0 max: 18 index: 5
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
	// min: -20 index: 5 max: 7 index: 4

}
