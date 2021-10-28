// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func rotationalCipher(input string, rotationFactor int) string {
	var out []byte
	out = make([]byte, len(input))

	// Write your code here
	for i := 0; i < len(input); i++ {
		switch {
		case input[i] >= 'a' && input[i] <= 'z':
			out[i] = input[i] + 3
			if out[i] > 'z' {
				out[i] = out[i]%'z' + 'a' - 1
			}

		case input[i] >= 'A' && input[i] <= 'Z':
			out[i] = input[i] + 3
			if out[i] > 'Z' {
				out[i] = out[i]%'Z' + 'A' - 1
			}

		case input[i] >= '0' && input[i] <= '9':
			out[i] = input[i] + 3

			if out[i] > '9' {
				out[i] = out[i]%'9' + '0' - 1
			}

		default:
			out[i] = input[i]
		}
	}

	return string(out)
}

func main() {
	// Call rotationalCipher() with test cases here
	out := rotationalCipher("Zebra-493?", 3)
	fmt.Println(out)
}
