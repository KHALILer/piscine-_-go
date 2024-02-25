package piscine

func AppendRange(min, max int) []int {
	array := []int{}
	for i := min; i < max; i++ {
		array = append(array, i+1)
	}
	return array
}
