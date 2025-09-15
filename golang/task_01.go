// Find the last non-repeating character in a given string. If all characters repeat, return an empty string.

// Example:

// > nonRepeat('candy canes do taste yummy')
// > 'u'

// ```

package main

// ```
import (
	"fmt"
	"strings"
)

var TEST_STRING string = "candy canes do taste yummy"

func nonRepeat(str string) string {
	var filteredMap map[string]int
	s := strings.Split(str, `''`)
	res := make(map[string]int)
	for val, ok := range s {
		if !ok {
			res[val] = 0
		}
		res[val] += 1
	}
	filteredMap()
	return res
}

func main() {
	fmt.Println(nonRepeat(TEST_STRING))
}
