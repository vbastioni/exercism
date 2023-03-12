package raindrops

import (
	"strconv"
	"strings"
)

// Convert a number to either a raindrop or itself as a string
func Convert(n int) string {
	sb := strings.Builder{}
	if n%3 == 0 {
		sb.WriteString("Pling")
	}
	if n%5 == 0 {
		sb.WriteString("Plang")
	}
	if n%7 == 0 {
		sb.WriteString("Plong")
	}
	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

var convs = []struct {
	n int
	w string
}{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

func ConvertArrayCompStr(n int) string {
	sb := strings.Builder{}
	for _, c := range convs {
		if n%c.n == 0 {
			sb.WriteString(c.w)
		}
	}
	if sb.String() == "" {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

func ConvertArrayCompLen(n int) string {
	sb := strings.Builder{}
	for _, c := range convs {
		if n%c.n == 0 {
			sb.WriteString(c.w)
		}
	}
	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

var convsM = map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}

func ConvertMapCompStr(n int) string {
	sb := strings.Builder{}
	for k, v := range convsM {
		if n%k == 0 {
			sb.WriteString(v)
		}
	}
	if sb.String() == "" {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

func ConvertMapCompLen(n int) string {
	sb := strings.Builder{}
	for k, v := range convsM {
		if n%k == 0 {
			sb.WriteString(v)
		}
	}
	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}
