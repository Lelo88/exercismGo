package isogram

import "strings"

func IsIsogram(word string) bool {
	letterMap := make(map[rune]bool)
	for _, letter := range strings.ToLower(word) {
		if letter == ' ' || letter == '-' {
			continue
		}
		if letterMap[letter] {
			return false
		}
		letterMap[letter] = true
	}
	return true
}
