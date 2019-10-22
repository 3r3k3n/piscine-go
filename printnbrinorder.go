package main

import "github.com/01-edu/z01"

func intToLetters(nb int) (letters []int) {
	for nb > 0 {
		if nb == 0 {
			letters = append(letters, 0)
		} else {
			letters = append(letters, nb%10)
		}
		nb = nb / 10
	}
	return
}
func sort(arr []int) []int {
	count := 0
	for range arr {
		count += 1
	}
	for i := count; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				value := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = value
			}
		}
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for _, c := range sort(intToLetters(nb)) {
		z01.PrintRune(rune(c) + '0')
	}
}
