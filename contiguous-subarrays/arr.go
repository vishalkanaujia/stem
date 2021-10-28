// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func countSubarrays(arr []int) []int {
	// Write your code here
	answer := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count := 1

		forward := i + 1
		backward := i - 1

		// forward subarrays
		for forward < len(arr) {
			if arr[i] > arr[forward] {
				count++
				forward++

				continue
			}

			break
		}

		for backward >= 0 {
			if arr[i] > arr[backward] {
				count++
				backward--

				continue
			}

			break
		}

		answer[i] = count
	}

	return answer
}

func main() {
	// Call countSubarrays() with test cases here
	answer := countSubarrays([]int{3, 4, 1, 6, 2})
	fmt.Printf("answer=%v\n", answer)
}
