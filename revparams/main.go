package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
		arguments := os.Args
		len := 0

		for lenght := range arguments {
			len++
		}

		for i := len; i < 0; i-- {
			wordRune := []rune(arguments[i])
			for _, words := range wordRune {
			z01.PrintRune(words)
		}
		z01.PrintRune(10)
	}
}
