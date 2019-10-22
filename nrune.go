package piscine

func NRune(s string, n int) rune {
	counterRune := []rune(s)
	counter := 0
	mainRune := 0
	for range s {
		counter++
	}
	if n > counter {
		return 0
	} else if n <= 0 {
		return '\x00'
	}
	for i := 1; i < n; i++ {
		mainRune = i
	}
	return counterRune[mainRune]
}
