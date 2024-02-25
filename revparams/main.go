package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		params := os.Args[i]
		for _, arg := range params {
			z01.PrintRune(arg)
		}
		z01.PrintRune('\n')
	}
}
