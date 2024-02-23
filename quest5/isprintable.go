package piscine

func IsPrintable(s string) bool {
	rn := []rune(s)
	f := false
	for i := range rn {
		if rn[i] >= 32 && rn[i] < 127 {
			f = true
		} else {
			return false
		}
	}
	return f
}
