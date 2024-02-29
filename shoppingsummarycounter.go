package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	Map := make(map[string]int)
	item := ""
	for _, words := range str {
		if words == ' ' {
			item += string(words)
		} else {
			if item != "" {
				Map[item]++
				item = ""
			}
		}
	}
	Map[item]++
	return Map
}
