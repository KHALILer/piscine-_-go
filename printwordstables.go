package piscine

import (
	"fmt"
)

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, char := range word {
			fmt.Print(char)
		}
	}
	fmt.Print('\n')
}
