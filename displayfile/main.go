package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	length := 0
	for i := range arguments {
		length = i + 1
	}

	if length > 1 {
		fmt.Print("Too many arguments")
		fmt.Println()

	} else if length == 0 {
		fmt.Print("File name missing")
		fmt.Println()

	} else {

		content, _ := ioutil.ReadFile(arguments[0])

		fmt.Printf(string(content))

	}
}
