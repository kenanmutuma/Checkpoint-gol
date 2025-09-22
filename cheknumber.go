package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	for _, s := range arg {
		if s >= '0' && s <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
