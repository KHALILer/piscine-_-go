package piscine

func Abort(a, b, c, d, e int) int {
	nb := []int{a, b, c, d, e}
	for i := 0; i < len(nb)-1; i++ {
		for j := 0; j < len(nb)-i-1; j++ {
			if nb[j] > nb[j+1] {
				nb[j], nb[j+1] = nb[j+1], nb[j]
			}
		}
	}
	return nb[len(nb)/2]
}
