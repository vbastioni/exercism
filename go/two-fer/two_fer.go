package twofer

import (
	"strings"
)

// ShareWith passes a cookie to the next person in the line then one to you
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	sb := strings.Builder{}
	sb.WriteString("One for ")
	sb.WriteString(name)
	sb.WriteString(", one for me.")
	return sb.String()
}
