package piscine

func Max(a []int) int {
	max := a[0]
	for _, i := range a {
		if max > i {
			return max
		}
	}
	return max
}
