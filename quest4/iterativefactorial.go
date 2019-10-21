package piscine

func IterativeFactorial(index int) int {
	result := 1
	if index < 0 || index > 50 {
		return 0
	} else {
		for i := 1; i <= index; i++ {
			result = result * i
		}
		return result
	}
}
