// Exercise 4.7:
// Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(s []byte) []byte {
	if utf8.RuneCount(s) == 1 {
		return s
	}
	_, size := utf8.DecodeRune(s)
	return append(reverseBytes(s[size:]), s[:size]...)
}

func main() {
	data := []byte("我爱你中国")
	fmt.Printf("%q\n", data)               // "我爱你中国"
	fmt.Printf("%q\n", reverseBytes(data)) // "国中你爱我"
}
