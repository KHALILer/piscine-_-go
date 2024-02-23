package piscine

func IsPrintable(s string) bool {
	rn := []rune(s)
	f := false
	for i := range rn {
		if rn[i] >= 'A' && rn[i] <= 'Z' || rn[i] >= 'a' && rn[i] <= 'z' {
			f = true
		} else {
			return false
		}
	}
	return f
}
