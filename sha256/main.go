package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//ii := []int{1}

	var m map[int]string
	m[1] = "11111"
	fmt.Printf("%t,%v", m == nil, m[1])
}
