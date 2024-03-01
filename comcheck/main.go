package main

import "os"

func main() {
	word := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	argument := os.Args[1:]
	for i := range argument {
		for _, s := range word {
			if argument[i] == s {
				count++
			}
		}
	}
	if count >= 1 {
		print("Alert!!!")
	}
	println()
}
