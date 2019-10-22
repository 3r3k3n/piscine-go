package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	yes := true
	for _, i := range runes {
		if !(i >= 'a' && i <= 'z') {
			yes = false
			return yes
		}
	}
	return yes
}