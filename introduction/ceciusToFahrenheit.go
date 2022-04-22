package main

import "fmt"

func celciusToFahrenheit() {
	var unit float64

	fmt.Print("Masukkan ukuran suhu pada fahrenheit: ")
	fmt.Scanln(&unit)

	fmt.Println("hasil convert ke dalam farenheit adalah:")
	fmt.Println(unit * 1.8 + 32)
}