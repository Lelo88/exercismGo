package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var action string
	switch log{
	case "❗ recommended search product 🔍", "❗ recommended product":
		action = "recommendation"
	case "executed search 🔍", "🔍 search recommended product ❗":
		action = "search"
	case "forecast: ☀ sunny", "☀ weather is sunny ❗":
		action = "weather"
	default:
		action = "default"
	}

	return action 
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
