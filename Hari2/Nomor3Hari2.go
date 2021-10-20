package main

import "fmt"

func main() {

	fmt.Println("=======================")
	fmt.Println("--------Nomor 3--------")
	fmt.Println("=======================")

	var bilangan, i int

	fmt.Println("Masukkan bilangan untuk ditemukan faktor bilangannya = ")
	fmt.Scanf("%d", &bilangan)
	fmt.Println("Faktor bilangan dari ", bilangan, " adalah =")
	for i = 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, "\t")
		}
	}
	fmt.Println()
}
