package isbn

func IsValidISBN(isbn string) bool {
	var sum, count int

    for index, character := range isbn {
        if character == '-' {
            continue
        }
        if character == 'X' && index == len(isbn)-1 {
            sum += 10
            count++
            break
        }
        if character < '0' || character > '9' {
            return false
        }

		// The character represents a digit. Convert it to an integer, this is done by subtracting the ASCII value of '0' from the ASCII value of the character.
		// The integer value of the character is then multiplied by the weight of the digit. The weight of the digit is 10 minus the position of the digit in the ISBN.
		// The sum of all the products is calculated.
        sum += int(character-'0') * (10 - count)
        count++
    }

    return count == 10 && sum%11 == 0
}
