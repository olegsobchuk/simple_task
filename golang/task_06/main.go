// Given an array of order objects for a restaurant, each with a table number and a list of ordered items, write a function that returns an object mapping each table number to a summary of how many times each item was ordered at that table. Extra credit: Could you go so far as to make this a restaurant management game?

// Example:

// ```
// const orders = [
//   { table: 1, items: ["burger", "fries"] },
//   { table: 2, items: ["burger", "burger", "fries"] },
//   { table: 1, items: ["salad"] },
//   { table: 2, items: ["fries"] }
// ];

// > orderSummary(orders)
// {
//   1: { burger: 1, fries: 1, salad: 1 },
//   2: { burger: 2, fries: 2 }
// }
// // or, string output format:
// "Table 1 ordered 1 burger, 1 fries, and 1 salad. Table 2 ordered 2 burgers and 2 fries."

// ```

package main

import "fmt"

type Order struct {
	table uint
	items []string
}

func orderSummary(orders []Order) map[uint]map[string]uint {
	output := map[uint]map[string]uint{}
	for _, order := range orders {
		_, ok := output[order.table]
		if !ok {
			output[order.table] = map[string]uint{}
		}

		lst := output[order.table]
		for _, item := range order.items {
			lst[item] += 1
		}
	}
	return output
}

func main() {
	orders := []Order{
		{table: 1, items: []string{"burger", "fries"}},
		{table: 2, items: []string{"burger", "burger", "fries"}},
		{table: 1, items: []string{"salad"}},
		{table: 2, items: []string{"fries"}},
	}
	fmt.Println(orderSummary(orders)) //
}
