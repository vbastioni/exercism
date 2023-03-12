package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type (
	File       []bool
	Chessboard map[string]File
)

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (sum int) {
	if f, ok := cb[file]; ok {
		for _, b := range f {
			if b {
				sum++
			}
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (sum int) {
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, f := range cb {
		if f[rank-1] {
			sum++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (sum int) {
	for _, f := range cb {
		for range f {
			sum++
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (sum int) {
	for _, f := range cb {
		for _, piece := range f {
			if piece {
				sum++
			}
		}
	}
	return
}
