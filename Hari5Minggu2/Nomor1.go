package main

import (
	"fmt"
	"sort"
)

// type FreqMap map[rune]int

// func StringFreq(s string) FreqMap {
// 	m := FreqMap{}
// 	for _, r := range s {
// 		m[r]++
// 	}
// 	return m
// }

// func Concurrent(strings string) {

// 	ch := make(chan FreqMap)

// 	go func(text string) {
// 		ch <- StringFreq(text)
// 	}(strings)
// 	for character, count := range <-ch {
// 		if string(character) != "," && string(character) != " " {
// 			fmt.Println(string(character), ":", count)
// 		}
// 	}
// 	return
// }

// func main() {

// 	Concurrent("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
// }

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	counter := make(map[string]int)
	for _, c := range text {
		//fmt.Printf("%v", string(c))
		if string(c) != " " {
			counter[string(c)]++
		}
	}

	fmt.Println(text)
	chars := make([]string, 0, len(counter))
	for chr := range counter {
		chars = append(chars, chr)
	}

	sort.Slice(chars, func(i, j int) bool {
		return counter[chars[i]] > counter[chars[j]]
	})
	for _, chr := range chars {
		fmt.Printf("%v: %v\n", chr, counter[chr])
	}

}
