package piscine

func ToUpper(s string) string {
	wordByte := []byte(s)
	for index, letters := range wordByte {
		if letters <= 122 && letters >= 97 {
			wordByte[index] = letters - 32
		}
	}
	return string(wordByte)
}
