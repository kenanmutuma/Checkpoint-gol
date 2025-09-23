package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	union := args[0] + args[1]
	result := []rune{}

	if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		return
	}

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
	for _, r := range result {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}

