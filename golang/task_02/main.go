// You’re assembling a custom mechanical keyboard. Each required part has a delivery time in days and an assembly time in hours. You can only assemble a part after it arrives, and you can only work on one part at a time. Given an array of parts where each part is { name, arrivalDays, assemblyHours }, return the minimum total hours needed to finish assembling all parts, starting from hour 0.

// Example:

// ```
// minAssemblyTime([
//   { name: "keycaps", arrivalDays: 1, assemblyHours: 2 },
//   { name: "switches", arrivalDays: 2, assemblyHours: 3 },
//   { name: "stabilizers", arrivalDays: 0, assemblyHours: 1 },
//   { name: "PCB", arrivalDays: 1, assemblyHours: 4 },
//   { name: "case", arrivalDays: 3, assemblyHours: 2 }
// ])

// > 74

// ```

package main

import (
	"fmt"
	"sort"
)

type keyboard struct {
	name          string
	arrivalDays   uint
	assemblyHours uint
}

var assemblyList []keyboard = []keyboard{
	{name: "keycaps", arrivalDays: 1, assemblyHours: 2},
	{name: "switches", arrivalDays: 2, assemblyHours: 3},
	{name: "stabilizers", arrivalDays: 0, assemblyHours: 1},
	{name: "PCB", arrivalDays: 1, assemblyHours: 4},
	{name: "case", arrivalDays: 3, assemblyHours: 2},
}

func minAssemblyTime(keyboards ...keyboard) uint {
	sort.Slice(keyboards, func(i, j int) bool {
		return keyboards[i].arrivalDays < keyboards[j].arrivalDays
	})
	lastKeyboard := keyboards[len(keyboards)-1]
	return lastKeyboard.arrivalDays*24 + lastKeyboard.assemblyHours
}

func main() {
	fmt.Println(minAssemblyTime(assemblyList...))
}
