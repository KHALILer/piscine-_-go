package piscine

func Unmatch(a []int) int {
	var count int
	for _, i := range a {
		count = 0
		for _, j := range a {
			if i == j {
				count++
			}
		}
		if count%2 != 0 {
			return i
		}
	}
	return -1
}
