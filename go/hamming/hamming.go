package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between two strands. Assuming
// letters does not change and stays in ASCII range.
func Distance(a, b string) (cnt int, err error) {
	if len(a) != len(b) {
		err = errors.New("strands length differs")
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return
}

// DistanceByRune calculates the Hamming Distance between two strands rune by rune
// and not char by char
func DistanceByRune(a, b string) (cnt int, err error) {
	ra, rb := []rune(a), []rune(b)
	if len(ra) != len(rb) {
		err = errors.New("strands length differs")
		return
	}
	for i := 0; i < len(ra); i++ {
		if ra[i] != rb[i] {
			cnt++
		}
	}
	return
}
