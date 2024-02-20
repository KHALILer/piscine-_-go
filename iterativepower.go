package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	} else {
		a := 1
		for i := 1; i <= power; i++ {
			a = nb
		}
		return a
	}
}
