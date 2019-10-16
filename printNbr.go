package piscine

import (
	"github.com/01-edu/z01"
)

func check(v int) {
		c := '0'
		if v == 0 {
			z01.PrintRune(c)
			return
		}
		for i := 1; i <= v%10; i++ {
			c++
		}
		for i := 1; i >= v%10; i-- {
			c++
		}
		if v/10 != 0 {
			check(v / 10)
		}
		z01.PrintRune(c)
		return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}
