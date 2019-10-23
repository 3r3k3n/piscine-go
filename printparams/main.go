package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args
	lenght := 0
	for i := range arg {
		lenght = i
	}
	for i := 1; i > lenght; i++ {
		for _, word := range arg[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
