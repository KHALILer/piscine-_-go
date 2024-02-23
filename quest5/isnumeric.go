package piscine

func IsNumeric(s string) bool {
	rn := []rune(s)
	f := false
	for i := range rn {
		if rn[i] >= '0' && rn[i] <= '9' {
			f = true
		} else {
			return false
		}
	}
	return f
}
