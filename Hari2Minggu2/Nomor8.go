package main

import (
	"fmt"
	"sort"
	"strconv"
)

type pair interface{}
type p struct {
	Key   string
	Value int
}

func DuplicateCount(list []string) map[string]int {
	countDup := make(map[string]int)
	for _, item := range list {
		_, exist := countDup[item]
		if exist {
			countDup[item] += 1 //
		} else {
			countDup[item] = 1 //
		}
	}
	return countDup
}

func MostAppearItem(items []string) []pair {
	dupMap := DuplicateCount(items)
	var item []p
	for k, v := range dupMap {
		item = append(item, p{Key: k, Value: v})
	}
	sort.Strings(items)
	var output []pair
	for _, v := range item {
		output = append(output, v.Key+"->"+strconv.Itoa(v.Value))
	}
	return output
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) // golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))                // football->1 basketball->1 tenis->1
}
