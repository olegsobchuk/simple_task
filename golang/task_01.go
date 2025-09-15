// Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

// Example:

// > nonRepeat('candy canes do taste yummy')
// > 'u'

// ```

package main

// ```
import (
	"fmt"
)

var TEST_STRING string = "candy canes do taste yummy"

func nonRepeat(str string) string {
	// Create a map to store character counts
	charCount := make(map[rune]int)

	// Count occurrences of each character
	for _, char := range str {
		charCount[char]++
	}

	// Find the last non-repeating character by iterating from the end
	for i := len(str) - 1; i >= 0; i-- {
		if charCount[rune(str[i])] == 1 {
			return string(str[i])
		}
	}

	// If no non-repeating character found, return empty string
	return ""
}

func main() {
	fmt.Println(nonRepeat(TEST_STRING))
}
