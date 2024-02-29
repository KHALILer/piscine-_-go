package piscine

func CollatzCountdown(start int) int {
	count := 0
	n := start
	if n > 0 {
		for n != 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
			count++
		}
		return count
	} else {
		return -1
	}
}
