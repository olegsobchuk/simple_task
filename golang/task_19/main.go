// Given a field represented as an array of 0s and 1s, where 1 means that position needs protection,
// you can place a scarecrow at any index, and each scarecrow protects up to k consecutive positions
// centered around itself (for example, for k = 3, a scarecrow at i protects i-1, i, and i+1). Return
// the minimum set of indices where scarecrows should be placed so that all the positions with 1 are
// protected. You can assume k is an odd number (or make up what happens if it's even).

// Examples:

// let field = [1, 1, 0, 1, 1, 0, 1];
// let k = 3;

// placeScarecrows(field, k);
// > [1, 4, 6]

// placeScarecrows([1, 0, 1, 1, 0, 1], k)
// > [1, 4]

// placeScarecrows([1, 1, 1, 1, 1], 1)
// > [0, 1, 2, 3, 4]

package main

import (
	"fmt"
)

func placeScarecrows(field []int, diameter int) (res []int) {
	shorts := [][]int{}
	for i := 0; i < len(field); i++ {
		if field[i] == 0 {
			continue
		}

		shorts = append(shorts, []int{i, i + diameter - 1})
		i = i + diameter - 1 // move to index after Scarecrows diameter
	}
	lastIdx := len(field) - 1
	for i := range shorts {
		point := min(shorts[i][0]+diameter/2, lastIdx)
		res = append(res, point)
	}

	return res
}

func main() {
	fmt.Println(placeScarecrows([]int{1, 1, 0, 1, 1, 0, 1}, 3)) // []int{1, 4, 6}
	fmt.Println(placeScarecrows([]int{1, 0, 1, 1, 0, 1}, 3))    //[]int{1, 4}
	fmt.Println(placeScarecrows([]int{1, 1, 1, 1, 1}, 1))       // []int{0, 1, 2, 3, 4}
}
