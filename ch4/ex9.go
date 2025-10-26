// Exercise 4.9:
// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileReader, err := os.Open("data.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	defer fileReader.Close()

	counts := make(map[string]int) // counts of Unicode words

	scanner := bufio.NewScanner(fileReader)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\tfreq\n")
	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
