package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var today = "Tuesday"

	switch today {
	case "Tuesday":
		fmt.Println("Cie baru selasa")
	case "Monday":
		fmt.Println("Baru juga senin")
	default:
		fmt.Println("Semangat bro bentar lagi")
	}
}
