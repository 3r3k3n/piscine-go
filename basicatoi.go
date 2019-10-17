package piscine

func intfor(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func BasicAtoi(s string) int {
	stroka := []rune(s)
	n := 0 
	for range stroka {
		n++
	}
	ans := 0
	for i := 0; i < n; i++ {
		if stroka[i] < '0' || stroka[i] > '9' {
			return ans
		} else {
			ans *=10
			ans += intfor(stroka[i])
		}
	}
	return ans

}