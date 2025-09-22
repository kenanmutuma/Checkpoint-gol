package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	res := []rune{}
	for i, save := 0, true; i < len(arg); save = !save {
		if save {
			end := i + num
			if end > len(arg) {
				end = len(arg)
			}
			for j := i; j < end; j++ {
				res = append(res, rune(arg[j]))
			}
		}
		i += num
	}
	return string(res)
}
