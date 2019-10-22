package piscine

func LastRune(s string) rune {
	Rune := []rune(s)
	end := 0
	for i := range Rune {
		end := i
	}
	return Rune[end]
}
