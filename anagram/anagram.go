package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	subjectCount := make(map[rune]int)
	for _, runeValue := range subject {
		subjectCount[runeValue]++
	}

	var results []string
candidateLoop:
	for _, candidate := range candidates {
		if strings.ToLower(candidate) == subject || len(candidate) != len(subject) {
			continue
		}

		candidateCount := make(map[rune]int)
		for _, r := range strings.ToLower(candidate) {
			candidateCount[r]++
		}

		for r, count := range subjectCount {
			if candidateCount[r] != count {
				continue candidateLoop
			}
		}

		results = append(results, candidate)
	}

	return results
}
