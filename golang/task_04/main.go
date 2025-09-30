// Given an array arr representing the positions of monsters along a straight line, and an integer d representing the minimum safe distance required between any two monsters, write a function to determine if all monsters are at least d units apart. If not, return the smallest distance found between any two monsters. If all monsters are safely spaced, return -1.

// Examples:

// ```
// let monsters = [3, 8, 10, 15];
// let d = 6;
// minMonsterDistance(bees, d)
// > 2

// minMonsterDistance([5, 9, 14, 18], 4)
// > -1

// ```

package main

import "fmt"

func minMonsterDistance(monsters []int, d int) int {
	minDist := -1
	if len(monsters) <= 1 {
		return minDist
	}

	for i := len(monsters) - 1; i > 0; i-- {
		dist := monsters[i] - monsters[i-1]
		if dist < d {
			if minDist == -1 || minDist > dist {
				minDist = dist
			}
		}
	}

	return minDist
}

func main() {
	monsters := []int{3, 8, 10, 15}
	d := 6
	fmt.Println(minMonsterDistance(monsters, d))            //> 2
	fmt.Println(minMonsterDistance([]int{5, 9, 14, 18}, 4)) //> -1
}
