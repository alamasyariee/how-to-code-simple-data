package main

import (
	"fmt"
	"math"
)

func sqrNumber() {
	var num float64
	
	fmt.Println("Masukkan angka yang kamu inginkan: ")
	fmt.Scanln(&num)
	
	fmt.Println("Akar dari angka yang anda masukkan adalah:")
	fmt.Println(math.Sqrt(num))

}