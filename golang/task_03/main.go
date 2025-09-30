// Given an array of audio file durations, write a function to group the files into playlists such that each playlist's total duration does not exceed a given limit maxDuration. Return an array of playlists, where each playlist is an array of file durations. Try to minimize the number of playlists.

// Example:

// ```
// const files = [120, 90, 60, 150, 80];
// const maxDuration = 200;

// groupAudioFiles(files, maxDuration)
// > [[150], [120,80], [90,60]]

// groupAudioFiles(files, 160)
// > [[150], [120], [90,60], [80]]

// ```

package main

import (
	"fmt"
	"slices"
	"sort"
)

func groupAudioFiles(files []uint, duration uint) [][]int {
	if len(files) == 0 {
		return [][]int{}
	}

	res := [][]int{}
	sort.Slice(files, func(i, j int) bool { return files[i] > files[j] })

	notUse := make([]bool, len(files))

	while slices.Contains(notUse, false) {
		sum := uint(0)
		group := []int{}
		for i, file := range files {
			if notUse[i] {
				continue
			}

		}
	}

	return res
}

func main() {
	files := []uint{120, 90, 60, 150, 80}
	maxDuration := uint(200)
	fmt.Println(groupAudioFiles(files, maxDuration)) // > [[150], [120,80], [90,60]]
	maxDuration = 160
	fmt.Println(groupAudioFiles(files, maxDuration)) // > [[150], [120], [90,60], [80]]
}
