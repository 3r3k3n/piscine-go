package piscine

import "github.com/01-edu/z01"

func Strlen(str string) {
	counter := 0
	aStringChangeable := []rune(str)
	for _, counter := range aStringChangeable {
		counter++
	} 
}
