package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := []rune{}

	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		tens := rune('0' + i/10)
		ones := rune('0' + i%10)
		result = append(result, tens, ones)

		if i == to {
			break
		}

		result = append(result, ',', ' ')
	}

	result = append(result, '\n')
	return string(result)
}
