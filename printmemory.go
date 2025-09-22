package main

import (
	"unicode"

	"github.com/01-edu/z01" // for PrintRune
)

func PrintMemory(arr [10]byte) {
	hexchar := "0123456789abcdef"

	for i := 0; i < len(arr); i += 4 {
		for j := 0; j < 4 && i+j < len(arr); j++ {

			b := arr[i+j]
			first := b / 16
			second := b % 16

			z01.PrintRune(rune(hexchar[first]))
			z01.PrintRune(rune(hexchar[second]))

			if j < 3 {
				z01.PrintRune(' ')
			}

		}
		z01.PrintRune('\n')
	}

	for i := 0; i < len(arr); i++ {
		r := rune(arr[i])
		switch {
		case unicode.IsPrint(r):
			z01.PrintRune(r)
		default:
			z01.PrintRune('.')

		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
