// Exercise 4.1:
// Write a function that counts the number of bits that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2)
package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i
var pc [256]uint8

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + uint8(i&0x1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// The Hamming Distance between c1 and c2 is 125
	fmt.Printf("The Hamming Distance between c1 and c2 is %d\n", hammingDistance(&c1, &c2))
}

func hammingDistance(h1, h2 *[32]byte) int {
	result := 0
	for i := 0; i < 32; i++ {
		result += int(pc[h1[i]^h2[i]])
	}
	return result
}
