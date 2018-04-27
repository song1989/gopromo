package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//rand.Int
	rnd, err := rand.Int(rand.Reader, big.NewInt(3))
	if err != nil {
		fmt.Printf("rand.Int() error : %v \\n", err)
	}
	fmt.Printf("rand.Int() : %v \\n", rnd)
}
