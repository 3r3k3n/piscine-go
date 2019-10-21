func interativFactorial(index int) int {

	result := 1
	if index > 0 {
		for i := 1; i <= index; i++ {
			result = result * i
		}
	} else {
		result = 0
	}
	return result
}

