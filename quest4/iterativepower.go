package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	} else {
		n := 1
		for i := 1; i <= power; i++ {
			n *= nb
		}
		return n
	}
}
