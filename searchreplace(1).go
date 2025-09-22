package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]

	if len(search) != 1 || len(replace) != 1 {
		return
	}

	searchRune := rune(search[0])
	replaceRune := rune(replace[0])

	result := ""

	for _, ch := range str {
		if ch == searchRune {
			result += string(replaceRune)
		} else {
			result += string(ch)
		}
	}
	fmt.Println(result)
}
