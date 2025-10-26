// Exercise 4.6:
// Write an in-place function that squashes each run of adjacent Unicode spaces(see unicode.IsSpace)
// in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
)

func squashSpaces(bytes []byte) []byte {
	out := bytes[:0]
	for _, b := range bytes {
		if unicode.IsSpace(rune(b)) {
			// remove leading spaces
			if len(out) != 0 && out[len(out)-1] != ' ' {
				out = append(out, ' ')
			}
		} else {
			out = append(out, b)
		}
	}
	return out
}

func main() {
	data := []byte("   \t\r我 爱  \t\r\n 中 \t\n国")
	fmt.Printf("%q\n", squashSpaces(data)) // "我 爱 中 国"
	fmt.Printf("%q\n", data)               // "我 爱 中 国\r\n 中 \t\n国"
}
