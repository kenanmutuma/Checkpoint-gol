package main

func RepeatAlpha(s string) string {
	result := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			count := int(ch-'a') + 1
			for i := 0; i < count; i++ {
				result += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			count := int(ch-'A') + 1
			for i := 0; i < count; i++ {
				result += string(ch)
			}
		} else {
			result += string(ch)
		}
	}

	return result
}
