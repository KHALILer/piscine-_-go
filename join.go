package piscine

func Join(strs []string, sep string) string {
	s := ""
	for i, word := range strs {
		s += word
		if i != len(strs)-1 {
			s += sep
		}
	}
	return s
}
