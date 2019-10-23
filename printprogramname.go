package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	for i := range arguments[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}