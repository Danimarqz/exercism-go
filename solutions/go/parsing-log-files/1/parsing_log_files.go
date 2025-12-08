package parsinglogfiles

import "regexp"


func IsValidLine(line string) bool {
	var valid = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return valid.MatchString(line)
}

func SplitLogLine(line string) []string {
	var sep = regexp.MustCompile(`<[-~*=]*>`)
	return sep.Split(line, -1)
}

func CountQuotedPasswords(lines []string) int {
	var re = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(line string) string {
	var reEOL = regexp.MustCompile(`end-of-line\d+`)
	return reEOL.ReplaceAllString(line, "")
}

func TagWithUserName(lines []string) []string {
	var userRe = regexp.MustCompile(`User\s+(\S+)`)
	out := make([]string, len(lines))

	for i, line := range lines {
		m := userRe.FindStringSubmatch(line)
		if len(m) == 2 {
			username := m[1]
			out[i] = "[USR] " + username + " " + line
		} else {
			out[i] = line
		}
	}
	return out
}