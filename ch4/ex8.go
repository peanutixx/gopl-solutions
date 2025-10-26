// Exercise 4.8:
// Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	CHAR_IS_SPACE = iota
	CHAR_IS_SYMBOL
	CHAR_IS_MARK
	CHAR_IS_DIGIT
	CHAR_IS_PRINT
	CHAR_IS_PUNCT
	CHAR_IS_LETTER
	CHAR_IS_NUMBER
	CHAR_IS_CONTROL
	CHAR_IS_GRAPHIC
)

func main() {
	counts := make(map[uint8]int)   // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encoding
	invalid := 0                    // count of invalid UTF-8 characters

	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsSpace(r):
			counts[CHAR_IS_SPACE]++
		case unicode.IsSymbol(r):
			counts[CHAR_IS_SYMBOL]++
		case unicode.IsMark(r):
			counts[CHAR_IS_MARK]++
		case unicode.IsDigit(r):
			counts[CHAR_IS_DIGIT]++
		case unicode.IsPrint(r):
			counts[CHAR_IS_PRINT]++
		case unicode.IsPunct(r):
			counts[CHAR_IS_PUNCT]++
		case unicode.IsLetter(r):
			counts[CHAR_IS_LETTER]++
		case unicode.IsNumber(r):
			counts[CHAR_IS_NUMBER]++
		case unicode.IsControl(r):
			counts[CHAR_IS_CONTROL]++
		case unicode.IsGraphic(r):
			counts[CHAR_IS_GRAPHIC]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	var category string
	for tag, count := range counts {
		switch tag {
		case CHAR_IS_SPACE:
			category = "SPACE"
		case CHAR_IS_SYMBOL:
			category = "SYMBOL"
		case CHAR_IS_MARK:
			category = "MARK"
		case CHAR_IS_DIGIT:
			category = "DIGIT"
		case CHAR_IS_PRINT:
			category = "PRINT"
		case CHAR_IS_PUNCT:
			category = "PUNCT"
		case CHAR_IS_LETTER:
			category = "LETTER"
		case CHAR_IS_NUMBER:
			category = "NUMBER"
		case CHAR_IS_CONTROL:
			category = "CONTROL"
		case CHAR_IS_GRAPHIC:
			category = "GRAPHIC"
		}
		fmt.Printf("%s\t%d\n", category, count)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, c := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, c)
		}
	}
	if invalid > 0 {
		fmt.Printf("%d invalid UTF-8 characters\n", invalid)
	}
}
