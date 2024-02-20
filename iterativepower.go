package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			nb *= nb
		}
		return nb
	}
}
