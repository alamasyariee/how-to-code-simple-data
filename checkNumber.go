package main

import "fmt"

func checkNumber(n int64) {
	if n == 0 {
		fmt.Println("Angka yang anda masukkan adalah nol")
	} else if n < 0 {
		fmt.Println("Angka yang anda masukkan adalah angka negatif")
	} else {
		fmt.Println("Angka yang anda masukkan adalah angka positif")
	}
}