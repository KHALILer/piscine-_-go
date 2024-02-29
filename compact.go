package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	N := 0

	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			s[N] = s[i]
			N++
		}
	}
	*ptr = s[:N]
	return N
}
