package main

import (
	"fmt"
)

func MostAppearItem(items []string) map[string]int {
	item := map[string]int{}

	for i := 0; i < len(items); i++ {
		if _, ok := item[items[i]]; ok {
			item[items[i]]++
		} else {
			item[items[i]] = 1
		}
	}

	return item
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))                // football->1 basketball->1 tenis->1
}
