package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			cnt++
		}
	}
	return cnt
}
