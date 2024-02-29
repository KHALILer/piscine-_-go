package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	Map := make(map[string]int)
	var item string
	for _, words := range str {
		if words == ' ' {
			Map[item]++
			item = ""
		} else if words != ' ' {
			item += string(words)
		}
	}
	Map[item]++
	return Map
}
