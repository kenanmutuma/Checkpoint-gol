package main

import (
	"fmt"
	"strings"
	//"strings"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	position := strings.Index(s, " ")
	if position == -1 {
		return s + "\n"
	}
	return s[:position] + "\n"
}
