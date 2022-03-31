package main

import "math/rand"


func RandomWithRange(min, max int) int {
// Body
var value = rand.Int() % (max - min + 1) + min

return value
}
