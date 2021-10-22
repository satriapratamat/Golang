package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	var newArrayA []string
	diffArray := true
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i] == arrayB[j] {
				diffArray = false
			}
		}
		if diffArray == true {
			newArrayA = append(newArrayA, arrayA[i])
		}
	}

	merge := append(newArrayA, arrayB[:]...)
	return merge

}

func main() {

	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// [“king”, “devil jin”, “akuma”, “eddie”, “steve”, “geese”]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// [“sergei”, “jin”, “steve”, “bryan”]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// [“alisa”, “yoshimitsu”, “devil jin”, “law”]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// [“devil jin”, “sergei”]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// [“hwoarang”]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}
