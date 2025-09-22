package main

import (
	"fmt"
)

func CountAlpha(s string) int {
	count := 0
	for _, ch := range s {
		if (ch <= 'z' && ch >= 'a') || (ch <= 'Z' && ch >= 'A') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
