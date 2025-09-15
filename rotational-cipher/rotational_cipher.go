package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var result string
	for i := 0; i < len(plain); i++ {
		char := plain[i]
		if char >= 'a' && char <= 'z' {
			result += string((char-'a'+byte(shiftKey))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+byte(shiftKey))%26 + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
