package piscine

func Rot14(s string) string {
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string((c-'a'+14)%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			result += string((c-'A'+14)%26 + 'A')
		} else {
			result += string(c)
		}
	}
	return result
}
