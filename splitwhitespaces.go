package piscine

func SplitWhiteSpaces(s string) []string {
	word := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if i > start {
				word = append(word, s[start:i])
			}
			start = i + 1
		} else if i == len(s)-1 {
			word = append(word, s[start:i+1])
		}
	}
	return word
}
