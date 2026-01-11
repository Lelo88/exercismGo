package wordcount

import "regexp"

type Frequency map[string]int

var (
	asciiLetterOrDigitPattern = regexp.MustCompile(`^[A-Za-z0-9]$`)
	asciiUppercasePattern     = regexp.MustCompile(`^[A-Z]$`)
)

// WordCount splits a phrase into normalized words and returns their occurrence counts.
func WordCount(phrase string) Frequency {
	frequencyMap := make(Frequency)
	currentWord := ""

	for index := 0; index <= len(phrase); index++ {
		var currentChar byte
		if index < len(phrase) {
			currentChar = phrase[index]
		} else {
			currentChar = ' '
		}

		if isLetterOrDigit(currentChar) {
			currentWord += string(toLowerASCII(currentChar))
		} else if currentChar == '\'' && shouldKeepApostrophe(phrase, index, currentWord) {
			currentWord += string(currentChar)
		} else {
			flushWord(frequencyMap, &currentWord)
		}
	}

	return frequencyMap
}

// isLetterOrDigit reports whether a byte is an ASCII letter or digit.
func isLetterOrDigit(character byte) bool {
	return asciiLetterOrDigitPattern.Match([]byte{character})
}

// toLowerASCII converts an uppercase ASCII letter to lowercase and leaves other bytes unchanged.
func toLowerASCII(character byte) byte {
	if asciiUppercasePattern.Match([]byte{character}) {
		return character + ('a' - 'A')
	}

	return character
}

// shouldKeepApostrophe reports whether an apostrophe should be kept inside the current word.
func shouldKeepApostrophe(phrase string, currentIndex int, currentWord string) bool {
	return currentIndex+1 < len(phrase) && isLetterOrDigit(phrase[currentIndex+1]) && len(currentWord) > 0
}

// flushWord stores the current word in the frequency map and resets the word buffer.
func flushWord(frequencyMap Frequency, currentWord *string) {
	if len(*currentWord) == 0 {
		return
	}

	frequencyMap[*currentWord]++
	*currentWord = ""
}
