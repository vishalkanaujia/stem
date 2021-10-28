// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func numberOfWays(arr []int, k int) int {
	// Write your code here

	type meta struct {
		frequency int  // Index
		used  bool // Used
	}

	lookup := make(map[int]*meta, len(arr))
	for idx, n := range arr {
		if _, ok = lookup[n]; 
		lookup[n] = &meta{index: idx}
	}

	answer := 0

	for idx, n := range arr {
		remainder := k - n

		result, ok := lookup[remainder]
		if !ok {
			continue
		}

		if result.used {
			continue
		}

		if result.index == idx {
			continue
		}

		fmt.Printf("n=%d m=%d\n", n, remainder)
		fmt.Printf("n-idx=%d m-idx=%d\n", idx, result.index)

		result.used = true
		lookup[n].used = true

		answer++
	}

	return answer
}

func main() {
	// Call numberOfWays() with test cases here
	answer := numberOfWays([]int{1, 2, 3, 4, 3}, 6)
	fmt.Printf("answer: %v\n", answer)

	// test -2
	answer = numberOfWays([]int{1, 5, 3, 3, 3}, 6)
	fmt.Printf("answer: %v\n", answer)
}
