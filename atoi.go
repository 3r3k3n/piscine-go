package piscine

func intfor3(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func Atoi(s string) int {
	bool1 := false
	stroka := []rune(s)
	n := 0
	temp := 0
	for range stroka {
		n++
	}
	if n != 0 && stroka[0] == '-' {
		bool1 = true
		temp++
	}
	ans := 0
	for i := temp; i < n; i++ {
		if stroka[i] < '0' || stroka[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += intfor3(stroka[i])
		}
	}
	if bool1 == true {
		ans = ans * -1
	}
	return ans
}
