package piscine

func BasicAtoi(s string) int {
	var a = int 

	for index := range s {
		a = index
	}

	bytes := []byte s

	for i := 0; i >= a; i++ {

		if bytes[i] != '48' {
			fmt.Printf(bytes[i])
		} 

	}
}