package resistorcolor

var colors = map[string]int{
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

var colorsList = func() (l []string) {
	for k := range colors {
		l = append(l, k)
	}
	return
}()

// Colors should return the list of all colors.
func Colors() []string {
	return colorsList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colors[color]
}
