package cipher

import (
	"regexp"
	"strings"
)

// To ensure each variation implements Cipher interface
var (
	_ Cipher = simple{}
	_ Cipher = shift(0)
	_ Cipher = vigenere("")
)

func getShifted(r, v rune) rune {
	if r >= 'A' && r <= 'Z' {
		r += 'a' - 'A'
	} else if r < 'a' || r > 'z' {
		return -1
	}
	r -= rune(v % 26)
	if r < 'a' {
		r += 26
	} else if r > 'z' {
		r -= 26
	}
	return r
}

type simple struct{}

func (s simple) Encode(input string) string {
	return strings.Map(func(r rune) rune { return getShifted(r, -('d' - 'a')) }, input)
}
func (s simple) Decode(input string) string {
	return strings.Map(func(r rune) rune { return getShifted(r, 'd'-'a') }, input)
}

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return simple{}
}

type shift int

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	return strings.Map(func(r rune) rune { return getShifted(r, -rune(c)) }, input)
}

func (c shift) Decode(input string) string {
	return strings.Map(func(r rune) rune { return getShifted(r, rune(c)) }, input)
}

type vigenere string

var reAllA = regexp.MustCompile(`^a*$`)

func validateKey(key string) bool {
	nk := strings.Map(func(r rune) rune {
		if r < 'a' || r > 'z' {
			return -1
		} else {
			return r
		}
	}, key)
	allA := reAllA.MatchString(key)
	return len(nk) != 0 && len(nk) == len(key) && !allA
}

func NewVigenere(key string) Cipher {
	if !validateKey(key) {
		return nil
	}
	return vigenere(key)
}

func (v vigenere) self() []rune {
	return []rune(v)
}

func (v vigenere) get(i int) rune {
	s := []rune(v)
	index := i % len(s)
	return s[index] - 'a'
}

func (v vigenere) Encode(input string) string {
	sb := strings.Builder{}
	lshift := 0
	for i, c := range input {
		if nc := getShifted(c, -v.get(i-lshift)); nc > -1 {
			sb.WriteRune(nc)
		} else {
			lshift++
		}
	}
	return sb.String()
}

func (v vigenere) Decode(input string) string {
	sb := strings.Builder{}
	lshift := 0
	for i, c := range input {
		if nc := getShifted(c, v.get(i-lshift)); nc > -1 {
			sb.WriteRune(nc)
		} else {
			lshift++
		}
	}
	return sb.String()
}
