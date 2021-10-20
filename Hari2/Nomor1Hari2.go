package main

import "fmt"

func main() {
	fmt.Println("=======================")
	fmt.Println("--------Nomor 1--------")
	fmt.Println("=======================")

	phi := 3.14
	T := 20.0
	r := 4.0

	luasAlas := phi * r * r
	luasSelimut := 2 * phi * r * T

	luasTabung := 2*luasAlas + luasSelimut

	fmt.Println("Luas tabung dengan jari-jari", r, "dan tinggi tabung", T, "adalah", luasTabung)

	fmt.Println("Sekarang, mari kita coba dengan cara input")
	fmt.Println("Masukkan panjang jari-jari lingkaran alas tabung: ")
	fmt.Scanln(&r)
	fmt.Println("Masukkan tinggi tabung: ")
	fmt.Scanln(&T)

	luasAlasInput := phi * r * r
	luasSelimutInput := 2 * phi * r * T

	luasTabungInput := 2*luasAlasInput + luasSelimutInput

	fmt.Println("Luas tabung dengan jari-jari", r, "dan tinggi tabung", T, "adalah", luasTabungInput)
}
