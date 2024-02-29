package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	var slice []int
	for i := max; i > min; i-- {
		slice = append(slice, i)
	}
	return slice
}
