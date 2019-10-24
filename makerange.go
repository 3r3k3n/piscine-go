package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	
	lenght := (max - min)
	answer := make([]int, lenght)
	 
	if min < max {
		for i := 0; i < lenght ; i++ {
			answer[i] = i + min
		}
	} 
	return answer 
}
