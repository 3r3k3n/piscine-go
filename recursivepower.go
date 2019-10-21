package piscine

func RecursivePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	} else {

		if power == 0 {
			return 1
		}
		if nb =! 1 {
			return RecursivePower(result, power-1) * result
		}

	}
	return result
}
