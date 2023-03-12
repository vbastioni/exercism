package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Score(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func BenchmarkScoreForEach(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreForEach(test.input)
		}
	}
}

// func BenchmarkScoreLower(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLower(test.input)
// 		}
// 	}
// }

// func BenchmarkScoreLowerForEach(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLowerForEach(test.input)
// 		}
// 	}
// }
// func BenchmarkScoreLowerManual(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLowerManual(test.input)
// 		}
// 	}
// }

// func BenchmarkScoreLowerManualForEach(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLowerManualForEach(test.input)
// 		}
// 	}
// }

// func BenchmarkScoreLowerUnicode(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLowerUnicode(test.input)
// 		}
// 	}
// }

// func BenchmarkScoreLowerUnicodeForEach(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("skipping benchmark in short mode.")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		for _, test := range scrabbleScoreTests {
// 			ScoreLowerUnicodeForEach(test.input)
// 		}
// 	}
// }
