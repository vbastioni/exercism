package resistorcolortrio

import (
	"fmt"
	"math"
)

var colors_ = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var units = []string{"ohms", "kiloohms", "megaohms", "gigaohms"}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

func Label(colors []string) string {
	n := (colors_[colors[0]]*10 + colors_[colors[1]]) * int(math.Pow10(colors_[colors[2]]))
	if n == 0 {
		return "0 ohms"
	}
	unit := int(math.Log10(float64(n)) / 3)
	return fmt.Sprintf("%d %s", n/int(math.Pow10(unit*3)), units[unit])
}
