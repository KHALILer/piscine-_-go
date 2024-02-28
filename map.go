package piscine

func Map(f func(int) bool, a []int) []bool {
	found := []bool{}
	for _, i := range a {
		found = append(found, f(i))
	}
	return found
}
