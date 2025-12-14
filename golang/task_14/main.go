// 	Write a function that determines if a number is
// [abundant, deficient, perfect, or amicable](https://www.encyclopedia.com/education/news-wires-white-papers-and-books/numbers-abundant-deficient-perfect-and-amicable?utm_source=cassidoo&utm_medium=email&utm_campaign=the-love-that-you-withhold-is-the-pain-that-you).

// Examples:

// whatKindOfNumber(6)
// > 'perfect'

// whatKindOfNumber(12)
// > 'abundant'

// whatKindOfNumber(4)
// > 'deficient'

package main

import (
	"fmt"
)

func kindOfNumber(num int) string {
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return numKind(sum, num)
}

func numKind(sum, num int) string {
	if sum == num {
		return "perfect"
	} else if sum > num {
		return "abundant"
	} else {
		return "deficient"
	}
}

func main() {
	fmt.Println("Task 13")
	fmt.Println(kindOfNumber(6))  // "perfect"
	fmt.Println(kindOfNumber(12)) // "abundant"
	fmt.Println(kindOfNumber(4))  // "deficient"
}
