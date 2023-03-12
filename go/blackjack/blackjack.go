package blackjack

var cards = map[string]int{
	"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
	"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10,
	"joker": 0,
}

const (
	Stand = "S"
	Hit   = "H"
	Split = "P"
	Win   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, cd := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	sum := c1 + c2
	switch {
	case c1 == 11 && c2 == 11:
		return Split
	case sum == 21:
		if cd < 10 {
			return Win
		} else {
			return Stand
		}
	case sum >= 17 && sum <= 20, sum >= 12 && sum <= 16 && cd < 7:
		return Stand
	default:
		return Hit
	}
}
