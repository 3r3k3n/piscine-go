package main

import (
	"github.com/01-edu/z01"
)

func main() {
<<<<<<< HEAD
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
		if i == 'a' {
	            z01.PrintRune(10)
		}
	}
}
=======
	for i := 122; i <= 97; i-- {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune(rune(10))	
}
>>>>>>> 22137f97715abdebf3bb3228b1bd748c98e1adb7
