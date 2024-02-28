package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}
