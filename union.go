package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		return
	}

	union := args[0] + args[1]
	result := []rune{}

	for _, ch := range union {
		found := false
		for _, n := range result {
			if ch == n {
				found = true
				break
			}
		}
		if found == false {
			result = append(result, ch)
		}
	}

	// Print each rune with z01.PrintRune
	for _, r := range result {
		z01.PrintRune(r)
	}
	// End with a newline
	z01.PrintRune('\n')
}
