// Exercise 4.5:
// Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func removeDuplicates(words []string) []string {
	i := 0
	for _, w := range words {
		if w != words[i] {
			words[i+1] = w
			i++
		}
	}
	return words[:i+1]
}

func removeDuplicates2(words []string) []string {
	result := words[:0]
	for _, w := range words {
		if len(result) == 0 || result[len(result)-1] != w {
			result = append(result, w)
		}
	}
	return result
}

func main() {
	data := []string{"one", "one", "two", "one", "two", "two", "three", "three", "three"}
	// fmt.Printf("%q\n", removeDuplicates(data)) 	   	// ["one" "two" "one" "two" "three"]
	fmt.Printf("%q\n", removeDuplicates2(data)) // ["one" "two" "one" "two" "three"]
	fmt.Printf("%q\n", data)                    // ["one" "two" "one" "two" "three" "two" "three" "three" "three"]
}
