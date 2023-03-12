package logs

import "strings"

var apps = map[rune]string{
	'☀': "weather",
	'🔍': "search",
	'❗': "recommendation",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		v, ok := apps[c]
		if ok {
			return v
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
