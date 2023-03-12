package scrabble

// BenchmarkScore-8                     	 7211684	       167.3 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreForEach-8              	 5761401	       208.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreLower-8                	 2335494	       512.2 ns/op	      16 B/op	       2 allocs/op
// BenchmarkScoreLowerForEach-8         	 2162811	       554.1 ns/op	      16 B/op	       2 allocs/op
// BenchmarkScoreLowerManual-8          	 4521302	       265.8 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreLowerManualForEach-8   	 4957197	       241.0 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreLowerUnicode-8          	 1426429	       840.2 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreLowerUnicodeForEach-8   	 2103975	       520.0 ns/op	       0 B/op	       0 allocs/op

import (
	"strings"
	"unicode"
)

func value(c byte) int {
	switch c {
	case 'D', 'G', 'd', 'g':
		return 2
	case 'B', 'C', 'M', 'P', 'b', 'c', 'm', 'p':
		return 3
	case 'F', 'H', 'V', 'W', 'Y', 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'K', 'k':
		return 5
	case 'J', 'X', 'j', 'x':
		return 8
	case 'Q', 'Z', 'q', 'z':
		return 10
	default:
		return 1
	}
}

func Score(word string) (sum int) {
	for i := 0; i < len(word); i++ {
		sum += value(word[i])
	}
	return
}

func ScoreForEach(word string) (sum int) {
	for _, c := range word {
		sum += value(byte(c))
	}
	return
}

func valueLower(c byte) int {
	switch c {
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 1
	}
}

func ScoreLower(word string) (sum int) {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		sum += valueLower(word[i])
	}
	return
}

func ScoreLowerForEach(word string) (sum int) {
	word = strings.ToLower(word)
	for _, c := range word {
		sum += valueLower(byte(c))
	}
	return
}

func valueToLower(c byte) int {
	if c < 'a' {
		c += ('a' - 'A')
	}
	switch c {
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 1
	}
}

func ScoreLowerManual(word string) (sum int) {
	for i := 0; i < len(word); i++ {
		sum += valueToLower(word[i])
	}
	return
}

func ScoreLowerManualForEach(word string) (sum int) {
	for _, c := range word {
		sum += valueToLower(byte(c))
	}
	return
}

func valueUnicode(c rune) int {
	c = unicode.ToLower(c)
	switch c {
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 1
	}
}

func ScoreLowerUnicode(word string) (sum int) {
	w := []rune(word)
	for i := 0; i < len(w); i++ {
		sum += valueUnicode(w[i])
	}
	return
}

func ScoreLowerUnicodeForEach(word string) (sum int) {
	for _, c := range word {
		sum += valueUnicode(c)
	}
	return
}
