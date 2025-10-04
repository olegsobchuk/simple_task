// Given a multi-line string and a sequence of Vim navigation commands (h for left, j for down, k for up, and l for right), and starting at the top-left character, write a function that processes the commands and returns the character under the cursor. If the cursor tries to move out of bounds, keep it at the last valid position.

// Example:

// ```
// const string = `Hello, world!
// how are ya?`; // or "Hello, world!\nhow are ya?"
// const commands = 'jlhll';

// getCharAfterVimCommands(string, commands)
// > 'w'

// ```

package main

import (
	"fmt"
	"strings"
)

func getCharAfterVimCommands(str, cmds string) string {
	lines := strings.Split(str, "\n")
	var cursorLine uint = 0
	var cursorPos uint = 0
	for _, cmd := range cmds {
		move(cmd, &cursorLine, &cursorPos)
	}
	line := lines[cursorLine]
	char := line[cursorPos]
	return string(char)
}

func move(r rune, nLine *uint, nPos *uint) {
	switch r {
	case 'h':
		if *nPos > 0 {
			*nPos -= 1
		}
	case 'l':
		*nPos += 1
	case 'k':
		if *nLine > 0 {
			*nLine -= 1
		}
	case 'j':
		*nLine += 1
	}

}

func main() {
	str := "Hello, world!\nhow are ya?"
	cmds := "jlhll"
	fmt.Println(getCharAfterVimCommands(str, cmds)) // "w"
}
