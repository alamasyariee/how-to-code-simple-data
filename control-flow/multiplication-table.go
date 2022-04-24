package main

import "fmt"

func multiplication() {
	var num int

	fmt.Print("Masukkan perkalian yang anda inginkan: ")
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "*", num, "=", i*num)
	}

}
