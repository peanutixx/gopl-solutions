// Exercise 4.2:
// Write a program that prints the SHA256 hash of its standard input by default
// but supports a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var length = flag.Int("n", 256, "Generate sha256/384/512 digest")

func main() {
	flag.Parse()

	// command-line argument validation
	algo := *length
	if algo != 256 && algo != 384 && algo != 512 {
		algo = 256
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch algo {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		}
	}
}
