package main

import "fmt"

func addTwoNuber() {
	var num1, num2 int
	
	fmt.Println("Massukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Println("Masukan angka kedua: ")
	fmt.Scanln((&num2))
	
	fmt.Println("Hasil dari penjumlahan kedua angka tersebut adalah:")
	fmt.Println(num1 + num2)

}