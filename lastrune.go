package piscine

func LastRune(s string) rune {
	rn := []rune(s)
  return rn[len(s)-1]
}
