// Given a CSV string where each row contains a name, age, and city (and values may be quoted,
// have embedded commas or escaped quotes), write a function that parses the CSV and outputs
// a formatted list of strings in the form: "Name, age Age, from City". Handle quoted fields
// containing commas and escaped quotes.

// Example:

// const csv = 'name,age,city\n"Ryu, Mi-yeong",30,"Seoul"\nZoey,24,"Burbank"'

// csvToList(csv)
// >
// - Ryu, Mi-yeong, age 30, from Seoul
// - Zoey, age 24, from Burbank

package main

import (
	"fmt"
	"strings"
)

func csvToList(str string) string {
	lines := strings.Split(str, "\n")
	formattedLines := []string{}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := parseCSVLine(line)
		s := fmt.Sprintf("- %s, age %s, from %s", fields[0], fields[1], fields[2])
		formattedLines = append(formattedLines, s)
	}
	return strings.Join(formattedLines, "\n")
}

func parseCSVLine(line string) []string {
	fields := []string{}
	currentField := ""
	inQuotes := false
	escaped := false

	for _, char := range line {
		switch char {
		case '"':
			if escaped {
				currentField += string(char)
				escaped = false
			} else {
				inQuotes = !inQuotes
			}
		case ',':
			if inQuotes {
				currentField += string(char)
			} else {
				fields = append(fields, strings.TrimSpace(currentField))
				currentField = ""
			}
		case '\\':
			if inQuotes {
				escaped = true
			} else {
				currentField += string(char)
			}
		default:
			currentField += string(char)
		}
	}
	fields = append(fields, strings.TrimSpace(currentField))
	return fields
}

func main() {
	fmt.Println("Task 17")

	csv := `name,age,city
	"Ryu, Mi-yeong",30,"Seoul"
	Zoey,24,"Burbank"`
	fmt.Println(csvToList(csv))
}
