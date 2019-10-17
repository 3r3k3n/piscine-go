package piscine 

import "github.com/01-edu/z01"

func PointOne(n *int) {
	*n = *n * 1
}

func piscine() {
	n := 0
	piscine.PointOne(&n)
	z01.Println(n)
}