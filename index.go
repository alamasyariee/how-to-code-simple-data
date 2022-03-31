package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	// var num int

	// fmt.Println(("Masukkan angka yang ingin diperiksa:"))
	// fmt.Scanln(&num)

	// IsPrime(num)

  rand.Seed(time.Now().Unix())
  var randomValue int

  randomValue = RandomWithRange(2, 10)
  fmt.Println("random number: ", randomValue);

  randomValue = RandomWithRange(2, 10)
  fmt.Println("random number: ", randomValue);

  randomValue = RandomWithRange(2, 10)
  fmt.Println("random number: ", randomValue);

}

