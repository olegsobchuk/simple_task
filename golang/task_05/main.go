// Write a generator function createLaundryItem() that returns an object representing a laundry item. This object should have a method nextCycle() which, when called, advances the item through a series of laundry cycles in order: "soak", "wash", "rinse", "spin", and "dry". After the final cycle, subsequent calls to nextCycle() should return "done".

// Example:

// ```
// let towel = createLaundryItem();

// console.log(towel.nextCycle()); // "soak"
// console.log(towel.nextCycle()); // "wash"
// console.log(towel.nextCycle()); // "rinse"
// console.log(towel.nextCycle()); // "spin"
// console.log(towel.nextCycle()); // "dry"
// console.log(towel.nextCycle()); // "done"
// console.log(towel.nextCycle()); // "done"

// ```

package main

import (
	"fmt"
)

var STATES = map[int]string{0: "init", 1: "soak", 2: "wash", 3: "rise", 4: "spin", 5: "dry", 6: "done"}

type LaundryItem struct {
	currentState int
}

func buildLaundryItem() LaundryItem {
	return LaundryItem{currentState: 0}
}

func (li *LaundryItem) nextCycle() string {
	tmpState := li.currentState + 1
	if STATES[tmpState] != "" {
		li.currentState = tmpState
	}
	return STATES[li.currentState]
}

func main() {
	towel := buildLaundryItem()

	fmt.Println(towel.nextCycle()) // "soak"
	fmt.Println(towel.nextCycle()) // "wash"
	fmt.Println(towel.nextCycle()) // "rinse"
	fmt.Println(towel.nextCycle()) // "spin"
	fmt.Println(towel.nextCycle()) // "dry"
	fmt.Println(towel.nextCycle()) // "done"
	fmt.Println(towel.nextCycle()) // "done"
}
