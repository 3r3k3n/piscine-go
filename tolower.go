package piscine

func ToLower(s string) string {
	wordByte := []byte(s)
	for index, letters := range wordByte {
		if letters <= 65 && letters >= 90 {
			wordByte[index] = letters + 32
		}
	}
	return string(wordByte)
}
