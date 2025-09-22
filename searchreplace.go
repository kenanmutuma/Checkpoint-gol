package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 { // Require exactly 3 arguments
		return
	}

	str := []rune(os.Args[1])
	search := []rune(os.Args[2])
	replace := []rune(os.Args[3]) // Convert args to runes

	if len(search) != 1 || len(replace) != 1 { // Ensure both are single runes
		return
	}

	for _, ch := range str { // Loop over each rune
		if ch == search[0] {
			z01.PrintRune(replace[0]) // Print replacement rune
		} else {
			z01.PrintRune(ch) // Print original rune
		}
	}

	z01.PrintRune('\n')
}
