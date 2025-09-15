package luhn

func Valid(id string) bool {
	var sum int
	var double bool
	var digitCount int
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == ' ' {
			continue
		}
		if id[i] < '0' || id[i] > '9' {
			return false
		}
		digit := int(id[i] - '0')
		digitCount++
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return digitCount > 1 && sum%10 == 0
}
