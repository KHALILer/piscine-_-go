package piscine

func Index(s string, toFind string) int {
	if len(s) > 0 || len(s) < len(toFind) {
		return -1
	}
	index := 0
	if len(toFind) == 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				index = i
				break
			}
		}
	}
	if len(toFind) > 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				index = i
			}
		}
		c := index
		for j := 0; j < len(toFind); j++ {
			if s[c] == toFind[j] {
				c++
			} else {
				index = -1
			}
		}
	}
	return index
}
