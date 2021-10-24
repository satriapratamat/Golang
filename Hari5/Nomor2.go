package main

import "fmt"

func caesar(offset int, input string) string {
	var inputRune = []rune(input)
	var newRune []rune
	if offset > 26 {
		offset = offset % 26
	}
	for _, value := range inputRune {
		// fmt.Println("value di line 16 = ", value)
		newValue := int(value) + offset
		// fmt.Println("newValue di line 14 = ", newValue)
		if newValue > 122 {
			// fmt.Println("newValue di line 16 = ", newValue)
			newValue = (newValue % 122) + 96
			// fmt.Println("newValue di line 18 = ", newValue)
		}
		newRune = append(newRune, rune(newValue))
	}
	return string(newRune)

}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	// bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
