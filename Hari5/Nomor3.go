package main

import "fmt"

func swap(a, b *int) {
	ubah := *a
	*a = *b
	*b = ubah
}

func main() {

	a := 10

	b := 20

	swap(&a, &b)

	fmt.Println("Nilai a yang baru =", a, "dan nilai b yang baru =", b)

}
