// Imagine a simplified version of the game Battleship played on a 2D grid. The grid represents the sea, and each
// cell can either be empty (.) or contain a part of a ship (X). Ships are placed horizontally or vertically, and there
// are no adjacent ships. Given a grid, count the number of battleships in it. Extra credit: can you make a layout
// generator for the game given these rules?

// Example:

// ```text
// const ships = [
//   ['X', 'X', '.', 'X'],
//   ['.', '.', '.', 'X'],
//   ['.', '.', '.', 'X'],
//   ['.', '.', '.', '.'],
// ];

// numberOfShips(ships)
// > 2

// ```

package main

import (
	"fmt"
	"slices"
)

func numberOfShips(field [][]rune) uint {
	var shipsCount uint = 0
	if len(field) == 0 {
		return 0
	}

	var ignors map[int][]int = make(map[int][]int)

	for rowIndex, row := range field {
		for colIndex, _ := range row {
			if slices.Contains(ignors[rowIndex], colIndex) {
				continue
			}
			if field[rowIndex][colIndex] == 'X' {
				shoot(&field, rowIndex, colIndex, ignors)
				shipsCount++
			}
		}
	}
	fmt.Println(ignors)
	return shipsCount
}

func shoot(field *[][]rune, row, col int, ignors map[int][]int) {
	if _, ok := ignors[row]; !ok {
		ignors[row] = []int{}
	}
	ignors[row] = append(ignors[row], col)
	if (*field)[row][col] == 'X' {
		if len((*field)[row])-1 > col {
			shoot(field, row, col+1, ignors)
		}
		if len(*field)-1 > row {
			shoot(field, row+1, col, ignors)
		}
	}
}

func main() {
	battleField := [][]rune{
		{'X', 'X', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', '.'},
	}
	fmt.Println(numberOfShips(battleField)) // Output: 2
}
