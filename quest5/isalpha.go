package piscine

func IsAlpha(s string) bool {
	rn := []rune(s)
	f := false
	for i := range rn {
		if rn[i] >= 'A' && rn[i] <= 'Z' || rn[i] >= 'a' && rn[i] <= 'z' || rn[i] >= '0' && rn[i] <= '9' {
			f = true
		} else {
			return false
		}
	}
	return f
}
