package main

import (
	"strconv"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	res := ""
	count := 1

	for i := 0; i < len(s); i++ {
		switch {
		case i+1 < len(s) && s[i] == s[i+1]:
			count++

		default:
			res += strconv.Itoa(count)
			res += string(s[i])
			count = 1
		}
	}
	return res
}
