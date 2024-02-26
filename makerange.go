package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	if min < 0 || max < 0 {
		return nil
	}

	array := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		array[i] = i + min
	}
	return array
}
