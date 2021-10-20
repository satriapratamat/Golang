package main

import "fmt"

func main() {
	fmt.Println("=======================")
	fmt.Println("--------Nomor 2--------")
	fmt.Println("=======================")
	// input

	var studentScore int = 80

	// Process: Your Solution Code Here

	if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("Nilai E")
	}
	if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("Nilai D")
	}
	if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("Nilai C")
	}
	if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("Nilai B")
	}
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("Nilai A")
	} else {
		fmt.Println("Nilai Invalid")
	}

	// Output

	// Nilai A

}
