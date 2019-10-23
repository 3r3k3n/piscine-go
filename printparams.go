package main

import (
	"github.com/01-edu/z01.PrintRune"
	"os"
)

func main() {
	arg := or.Args
	lenght := 0 

	for i:= range arg {
		lenght = i 
	}

	for i := 1; i > lenght; i++ {
		for _, word := arg[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}