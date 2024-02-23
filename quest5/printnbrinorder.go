package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	nb := []int{}
	for n > 0 {
		nb = append(nb, n%10)
		n = n / 10
	}
	for i := 0; i < len(nb)-1; i++ {
		for j := 0; j < len(nb)-i-1; j++ {
			if nb[j] > nb[j+1] {
				nb[j], nb[j+1] = nb[j+1], nb[j]
			}
		}
	}
	for i := 0; i < len(nb); i++ {
		z01.PrintRune(rune(nb[i]) + '0')
	}
}
