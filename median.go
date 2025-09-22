package main

import (
	"fmt"
	"sort"
)

func Median(digits []int) float64 {
	if len(digits) == 0 {
		return 0
	}

	sort.Ints(digits)

	n := len(digits)
	mid := n / 2

	if n%2 == 1 {
		return float64(digits[mid])
	}
	return float64(digits[mid-1]+digits[mid]) / 2.0
}

func main() {
	digits:= []int{10, 8, 5, 7, 2}
	fmt.Println(Median(digits))
	fmt.Println(digits)
}       
