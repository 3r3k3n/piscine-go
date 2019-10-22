package piscine

func ToLower(s string) string {
	wordByte := []byte(s)
	for index, letters := range wordByte {
		if letters <= 90 && letters >= 65 {
			wordByte[index] = letters + 32
		}
	}
	return string(wordByte)
}
