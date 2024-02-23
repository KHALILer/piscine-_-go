package piscine

func IsUpper(s string) bool {
	rn := []rune(s)
	f := false
	for i := range rn {
		if rn[i] >= 'A' && rn[i] <= 'Z' {
			f = true
		} else {
			return false
		}
	}
	return f
}
