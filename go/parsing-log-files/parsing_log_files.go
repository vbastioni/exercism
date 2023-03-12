package parsinglogfiles

import "regexp"

var (
	reValidLine   = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	rePrettySplit = regexp.MustCompile(`<[~\-\*=]*>`)
	reQuotedPass  = regexp.MustCompile(`(?i)".*?password.*?"`)
	reRemEOL      = regexp.MustCompile(`end-of-line\d*`)
	reReplUser    = regexp.MustCompile(`^(.*(?:User\s+(\w+)).*)$`)

	reReplUserSubStr = "[USR] $2 $1"
)

func IsValidLine(text string) bool {
	return reValidLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return rePrettySplit.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (sum int) {
	for _, l := range lines {
		if reQuotedPass.MatchString(l) {
			sum++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	return reRemEOL.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (amended []string) {
	for _, line := range lines {
		amended = append(amended, reReplUser.ReplaceAllString(line, reReplUserSubStr))
	}
	return
}
