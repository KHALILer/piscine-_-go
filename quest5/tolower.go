package piscine

func ToLower(s string) string {
	rn := []rune(s)
	for i := 0; i < len(s); i++ {
		if rn[i] >= 'A' && rn[i] <= 'Z' {
			rn[i] += 32
		}
	}
	return string(rn)
}
