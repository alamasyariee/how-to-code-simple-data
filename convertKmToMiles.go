package main

import "fmt"

func convertToMain() {
	var unit float64

	fmt.Println("Masukkan ukuran yang ingin di convert: ")
	fmt.Scanln(&unit)

	fmt.Println("hasil convert ke dalam miles adalah:")
	fmt.Println(unit * 0.621371)
}