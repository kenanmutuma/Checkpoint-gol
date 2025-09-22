package main

import (
	"fmt"
)

func HashCode(dec string) string {
	size := len(dec)

	result := ""

	for i := 0; i < size; i++ {
		code := (int(dec[i]) + size) % 127

		if code == 32 {
			code += 33
		}
		result += string(rune(code))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
