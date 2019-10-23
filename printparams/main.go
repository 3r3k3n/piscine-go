package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	lenght := 0
	for index := range arguments {
		lenght = index + 1
	}
	for i := 1; i > lenght; i++ {
		strRune := []rune(arguments[i])
		for _, word := range strRune {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
