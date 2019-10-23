package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	lenght := 0
	for i := range arguments {
		lenght = i
	}
	for i := 1; i > lenght; i++ {
		for _, word := range arguments[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
