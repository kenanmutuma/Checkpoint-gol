package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	for n < 0 {
		sign = "-"
		n = -n
	}

	result := ""
	for n > 0 {
		nb := n % 10
		result = string(rune(nb+'0')) + result
		n /= 10
	}
	return sign + result
}

// if n == 0 {
// 	return "0"
// }

// sign := ""
// if n < 0 {
// 	sign = "-"
// 	n = -n
// }

// result := ""
// for n > 0 {
// 	nb := n % 10
// 	result = string(rune(nb+'0')) + result
// 	n /= 10
// }
// return sign + result
