package main

import "fmt"

func IsPrime(n int) {
	isTrue := true

	if n == 1 {
		fmt.Printf("Jangan masukkan angka satu \n")
	}

	for i := 2; i < n; i++ {

		if (n % i == 0) {
			isTrue = false
			break
		}
	}

	if (isTrue) {
		fmt.Println("Angka yang anda masukkan adalah bilangan prima")
	} else {
		fmt.Println("Angka yang anda masukkan bukanlah bilangan prima")
	}

}
