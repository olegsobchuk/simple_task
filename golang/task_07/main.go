// Given an array of side lengths, write a function to determine they can form a hexagon
// with three side-length pairs (as in, three pairs of equal sides needed). Return true if possible.

// Examples:

// ```
// canFormHexagon([2, 3, 8, 8, 2, 3])
// > true

// canFormHexagon([1, 2, 3, 4, 5, 6])
// > false

// canFormHexagon([2, 2, 2, 2, 2, 2, 2])
// > false

// ```

package main

import (
	"fmt"
	"slices"
)

func canFormHexagon(sides []int) bool {
	if len(sides) != 6 {
		return false
	}
	slices.Sort(sides)
	for i := range sides {
		if i%2 == 1 {
			continue
		}
		if sides[i] != sides[i+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canFormHexagon([]int{2, 3, 8, 8, 2, 3}))    //true
	fmt.Println(canFormHexagon([]int{1, 2, 3, 4, 5, 6}))    //false
	fmt.Println(canFormHexagon([]int{2, 2, 2, 2, 2, 2, 2})) //false
}
