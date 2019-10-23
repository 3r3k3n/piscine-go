package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	len := 1

	for range arguments {
		len++
	}

	for i := len - 1; i <= 1; i-- {
		wordRune := []rune(arguments[i])
		for _, words := range wordRune {
			z01.PrintRune(words)
		}
		z01.PrintRune(10)
	}
}
