package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHead []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHead)

	var count int
	output := 0
	for count < len(dragonHead) {
		for i := 0; i < len(knightHead); i++ {
			if knightHead[i] >= dragonHead[count] {
				output += knightHead[i]
				count++
				break
			}
			if i == len(knightHead)-1 && i != len(dragonHead) {
				fmt.Println("knight fall")
				return
			}
		}
	}
	fmt.Println(output)
}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
