// Given an array of strings representing the names of monarchs and their ordinal numbers,
// write a function that returns the list of names sorted first by name and then
// by their ordinal value (in Roman numerals), in ascending order.

// Example:

// > sortMonarchs(["Louis IX", "Louis VIII", "Philip II", "Philip I"])
// > ["Louis VIII", "Louis IX", "Philip I", "Philip II"]

// > sortMonarchs(["George VI", "George V", "Elizabeth II", "Edward VIII"])
// > ["Edward VIII", "Elizabeth II", "George V", "George VI"]

package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func sortMonarchs(names []string) []string {
	slices.SortFunc(names, func(a, b string) int {
		nameA := splitName(a)
		nameB := splitName(b)
		return cmp.Or(
			cmp.Compare(nameA[0], nameB[0]),
			cmp.Compare(decodeRome(nameA[1]), decodeRome(nameB[1])),
		)
	})
	return names
}

func splitName(title string) []string {
	ind := strings.LastIndex(title, " ")
	name, rank := title[:ind], title[ind+1:]
	return []string{name, rank}
}

func decodeRome(roman string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var decNum, tmpNum int
	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]
		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	return decNum
}

func main() {
	fmt.Println(sortMonarchs([]string{"Louis IX", "Louis VIII", "Philip II", "Philip I"}))
	// ["Louis VIII", "Louis IX", "Philip I", "Philip II"]
	fmt.Println(sortMonarchs([]string{"George VI", "George V", "Elizabeth II", "Edward VIII"}))
	//  ["Edward VIII", "Elizabeth II", "George V", "George VI"]
}
