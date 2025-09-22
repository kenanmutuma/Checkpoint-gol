package main

import (
	"fmt"
	"os"
)

func piscine() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	s := os.Args[1]
	res := ""
	spaceNeeded := false
	isword := false

	for _, ch := range s {
		if ch != ' ' && ch != '\t' {
			if spaceNeeded {
				res += " "
				spaceNeeded = false
			}
			res += string(ch)
			isword = true

		} else if isword {
			spaceNeeded = true
			isword = false
		}
	}
	fmt.Println(res)
}
