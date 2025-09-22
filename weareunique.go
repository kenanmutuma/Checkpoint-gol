package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}

	count := 0
	seen := ""

	for _, c := range str1 + str2 {
		if !contains(seen, c) {
			if !contains(str1, c) || !contains(str2, c) {
				count++
			}
			seen += string(c)
		}
	}
	return count
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
