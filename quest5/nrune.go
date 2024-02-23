package piscine

func NRune(s string, n int) rune {
	if len(s) >= n && len(s) != 0 && n > 0 {
		rn := []rune(s)
		return rn[n-1]
	} else {
		return 0
	}
}
