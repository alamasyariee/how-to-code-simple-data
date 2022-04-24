package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("Masukkan angka yang anda inginkan: ")
	fmt.Scanln(&num1)

	fmt.Println("Masukkan angka yang anda inginkan: ")
	fmt.Scanln(&num2)

	fmt.Println("Masukkan angka yang anda inginkan: ")
	fmt.Scanln(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println("Dari ketiga angka terinput, yang terbesar adalah angka", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("Dari ketiga angka terinput, yang terbesar adalah angka", num2)
	} else {
		fmt.Println("Dari ketiga angka terinput, yang terbesar adalah angka", num3)
	}

}
