package bob

import (
    "strings"
    "unicode"
)

// Hey returns Bob's response to a remark.
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)

    if remark == "" {
        return "Fine. Be that way!"
    }

    isQuestion := strings.HasSuffix(remark, "?")
    isYelling := isUpper(remark)

    switch {
    case isYelling && isQuestion:
        return "Calm down, I know what I'm doing!"
    case isQuestion:
        return "Sure."
    case isYelling:
        return "Whoa, chill out!"
    default:
        return "Whatever."
    }
}

// isUpper checks if the remark is in uppercase.
func isUpper(remark string) bool {
    hasLetters := false
    for _, character := range remark {
        if unicode.IsLetter(character) {
            hasLetters = true
            if unicode.IsLower(character) {
                return false
            }
        }
    }
    return hasLetters
}