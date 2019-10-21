package main

import "github.com/01-edu/z01"

func Raid1ba(x, y int) {
	if x < 0 || y < 0 {
	} else {
		if x == 0 || y == 0 {
		} else {
			if x == 1 && y == 1 {
				z01.PrintRune('/')
				z01.PrintRune(10)
			} else {
				if y == 1 {
					for i := 1; i <= x; i++ {
						if i == 1 {
							z01.PrintRune('/')
						} else if i == x {
							z01.PrintRune('\\')
						} else {
							z01.PrintRune('*')
						}
					}
					z01.PrintRune(10)
				} else {
					for i := 1; i <= x; i++ {
						if i == 1 {
							z01.PrintRune('/')
						} else if i == x {
							z01.PrintRune('\\')
						} else {
							z01.PrintRune('*')
						}
					}
					z01.PrintRune(10)
					for j := 1; j <= y-2; j++ {
						for k := 1; k <= x; k++ {
							if k == 1 || k == x {
								z01.PrintRune('*')
							} else {
								z01.PrintRune(' ')
							}
						}
						z01.PrintRune(10)
					}
					for i := 1; i <= x; i++ {
						if i == 1 {
							z01.PrintRune('\\')
						} else if i == x {
							z01.PrintRune('/')
						} else {
							z01.PrintRune('*')
						}
					}
					z01.PrintRune(10)
				}
			}
		}
	
	}
	func main() {
		student.Raid1a(5,3)
	}
}
