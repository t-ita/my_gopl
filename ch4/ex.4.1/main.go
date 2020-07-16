package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] は i のポピュレーションカウントです
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDiff(c1, c2 [32]byte) int {
	var cnt int

	for i := 0; i < 32; i++ {
		cnt += int(pc[c1[i]^c2[i]])
	}

	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println("Diff Count = ", countDiff(c1, c2))
}
