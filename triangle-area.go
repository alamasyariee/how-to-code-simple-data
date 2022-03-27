package main

import "fmt"

func triangleArea() {
	var base, height float32
	
	fmt.Println("Masukkan ukuran alas segitiga: ")
	fmt.Scanln(&base)

	fmt.Println("Masukkan ukuran tinggi segitiga: ")
	fmt.Scanln(&height)
	
	fmt.Println("Luas segitiga adalah:")

	fmt.Println((base * height) / 2)

}