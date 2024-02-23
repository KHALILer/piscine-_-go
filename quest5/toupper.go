package piscine

func ToUpper(s string) string {
	rn := []rune(s)
	for i := 0; i < len(s); i++ {
		if rn[i] >= 'a' && rn[i] <= 'z' {
			rn[i] -= 32
		}
	}
	return string(rn)
}
