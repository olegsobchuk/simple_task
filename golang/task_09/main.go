// Given an array of strings representing a sequence of traffic light states ("red" for stop, "green" for go, "yellow"
// for slow), write a function that returns true if the sequence could represent a valid state machine for a standard
// traffic light. The only valid transitions are: red to green, green to yellow, and yellow to red.

// Example:

// ```text
// > isValidTrafficSequence(["red", "green", "yellow", "red", "green"])
// > true

// > isValidTrafficSequence(["red", "yellow", "green"]);
// > false

// > isValidTrafficSequence([])
// > true

// ```

package main

import (
	"fmt"
)

var STATES = map[string]string{"red": "green", "green": "yellow", "yellow": "red"}

func isValidTrafficSequence(list []string) bool {
	var prev string
	for i, val := range list {
		if i == 0 {
			prev = val
			continue
		}
		if STATES[prev] != val {
			return false
		}
		prev = val
	}

	return true
}

func main() {
	fmt.Println(isValidTrafficSequence([]string{"red", "green", "yellow", "red", "green"})) // true
	fmt.Println(isValidTrafficSequence([]string{"red", "yellow", "green"}))                 // false
	fmt.Println(isValidTrafficSequence([]string{}))                                         // true
}
