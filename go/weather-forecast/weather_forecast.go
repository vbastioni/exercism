// Package weather gives the forecasted weather for the current location and condition
// in the country.
package weather

// CurrentCondition holds information about the current weather condition,
// if it's raining or if it's sunny.
var CurrentCondition string

// CurrentLocation holds information about the location of the weather condition.
var CurrentLocation string

// Forecast update the weather conditions for a location and returns a
// beautifully formatted string about it.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
