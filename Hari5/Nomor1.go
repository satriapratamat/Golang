package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	if strings.Contains(b, a) {
		return a
	}
	if strings.Contains(a, b) {
		return b
	} else {
		return "Kalimat a tidak mengandung kalimat b atau sebaliknya"
	}

}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA

}
