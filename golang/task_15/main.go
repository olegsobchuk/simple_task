// Given the non-negative integer n , output the value of its
// [hyperfactorial](https://mathworld.wolfram.com/Hyperfactorial.html?utm_source=cassidoo&utm_medium=email&utm_campaign=i-recommend-the-freedom-that-comes-from-asking).
// Don't worry about outputs exceeding your language's integer limit.

// Examples:

// > hyperfactorial(0)
// > 1

// > hyperfactorial(2)
// > 4
// >
// > hyperfactorial(3)
// > 108

// > hyperfactorial(7)
// > 3319766398771200000

package main

import (
	"fmt"
	"math"
)

func hyperfactorial(n int) int {
	sum := 1

	for i := 1; i <= n; i++ {
		sum *= int(math.Pow(float64(i), float64(i)))
	}
	return sum
}

func main() {
	fmt.Println("Task 15")
	fmt.Println(hyperfactorial(0)) // 1
	fmt.Println(hyperfactorial(2)) // 4
	fmt.Println(hyperfactorial(3)) // 108
	fmt.Println(hyperfactorial(7)) // 3319766398771200000
}
