package main

import "fmt"

func printPrimeNumber() {
	var lower, high int

	fmt.Print("Masukkan batas nilai terendah: ")
	fmt.Scanln(&lower)

	fmt.Print("Masukkan batas nilai tertinggi: ")
	fmt.Scanln(&high)

	if lower < 2 {
		fmt.Println("Angka terendah tidak boleh kurang dari 2")
	} else {
		for i := lower; i <= high; i++ {
			if isPrime(i) {
				fmt.Println(i)
			}
		}
	}

}

func isPrime(n int) bool {
	isTrue := true

	if n == 1 {
		fmt.Printf("Jangan masukkan angka satu \n")
	}

	for i := 2; i < n; i++ {

		if n%i == 0 {
			isTrue = false
			break
		}
	}

	return isTrue

}
