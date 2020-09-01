package main

import (
	"fmt"
	"math/rand"
)

func generateRand() int {
	num := rand.Intn(100)
	return num
}

func main() {
	// 1. generate a random Inter
	target := generateRand()
	fmt.Println(target)
}
