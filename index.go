package main

import "fmt"


func main() {

	var num int

	fmt.Println(("Masukkan angka yang ingin diperiksa:"))
	fmt.Scanln(&num)
	IsPrime(num)
}