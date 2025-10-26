// Exercise 4.4:
// Write a version of rotate that operates in a single pass.
package main

import "fmt"

func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	moved := 0
	for start := 0; moved < n; start++ {
		curr := start
		value := nums[curr] // the value to move
		for {
			next := (curr + k) % n
			nums[next], value = value, nums[next]
			moved++
			curr = next
			if curr == start {
				break
			}
		}
	}
}

func rotateLeft(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	moved := 0
	for start := 0; moved < n; start++ {
		curr := start
		value := nums[curr] // the value to move
		for {
			next := (curr + n - k) % n
			nums[next], value = value, nums[next]
			moved++
			curr = next
			if curr == start {
				break
			}
		}
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	// rotateRight(a, 2)
	rotateLeft(a, 2)
	fmt.Println(a)
}
