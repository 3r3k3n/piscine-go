package piscine

func AlphaCount(str string) int {
	lengh := 0
	result := 0
	for range str {
		lengh++
	}
	for i := 0; i < lengh; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' {
			result++
		}
	}
	return result
}
