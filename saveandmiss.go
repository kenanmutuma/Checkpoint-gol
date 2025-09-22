package main

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	runes := []rune(arg)
	result := ""
	save := true

	for i := 0; i < len(runes); i += num {
		end := i + num
		if end > len(runes) {
			end = len(runes)
		}

		if save {
			result += string(runes[i:end])
		}
		save = !save
	}
	return result
}
