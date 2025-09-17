// Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

// Example:

// > nonRepeat('candy canes do taste yummy')
// > 'u'

package main

import (
	"fmt"
)

var TEST_STRING string = "candy canes do taste yummy"

func nonRepeat(str string) string {
	charCount := make(map[rune]uint)
	runes := []rune(str)

	for _, chr := range runes {
		charCount[chr]++
	}

	for n := len(runes) - 1; n >= 0; n-- {
		element := runes[n]
		if charCount[element] == 1 {
			return string(element)
		}
	}
	return ""
}

func main() {
	fmt.Println(nonRepeat(TEST_STRING))
}
