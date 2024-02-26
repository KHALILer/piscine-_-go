package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	array := make([]int, max-min+1)

	for i := 0; i < max-min+1; i++ {
		array[i] = i + min
	}
	return array
}
