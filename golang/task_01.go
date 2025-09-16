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
	arr := make(map[rune]uint)
	for _, chr := range str {
		arr[chr]++
	}
	for n := len(str) - 1; n > 0; n-- {
		element := rune(str[n])
		if arr[element] == 1 {
			return string(element)
		}
	}
	return ""
}

func main() {
	fmt.Println(nonRepeat(TEST_STRING))
}
