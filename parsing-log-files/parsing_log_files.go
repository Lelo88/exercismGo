package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(ERR|INF|TRC|DBG|WRN|FTL)\] `).MatchString(text)
}

func SplitLogLine(text string) []string {
	separatorRegex := regexp.MustCompile(`<[-~*=]*>`)
    fields := separatorRegex.Split(text, -1)
    return fields
}

func CountQuotedPasswords(lines []string) int {
	passwordRegex, err := regexp.Compile(`"(.*password.*)"`)

	if err != nil {
		panic(err)
	}

	count := 0
	for _, line := range lines {
		count += len(passwordRegex.FindStringSubmatch(line))
	}
	
	return count
}

func RemoveEndOfLineText(text string) string {
	removeRegex := regexp.MustCompile(`end-of-line\d+`)
	return removeRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	userRegex := regexp.MustCompile(`User\s+(\S+)`)
    var result []string

    for _, line := range lines {
        matches := userRegex.FindStringSubmatch(line)
        if len(matches) > 1 {
            userName := matches[1]
            taggedLine := "[USR] " + userName + " " + line
            result = append(result, taggedLine)
        } else {
            result = append(result, line)
        }
    }

    return result
}
