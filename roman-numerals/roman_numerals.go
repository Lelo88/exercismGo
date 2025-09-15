package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
        return "", errors.New("out of range")
    }

    romanNumerals := []struct {
        Value int
        Symbol string
    }{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    result := ""
    for _, mappedDigit := range romanNumerals {
        for input >= mappedDigit.Value {
            result += mappedDigit.Symbol
            input -= mappedDigit.Value
        }
    }

    return result, nil
	
}
