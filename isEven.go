package main

import "fmt"

func IsEven(n int64) {
	if n % 2 == 0 {
		fmt.Println("Angka yang anda masukkan adalah angka genap")
	} else {
		fmt.Println("Angka yang anda masukkan adalah angka ganjil")
	}
}