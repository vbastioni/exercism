package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

var tus = map[TemperatureUnit]string{Celsius: "°C", Fahrenheit: "°F"}

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	return tus[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string { return fmt.Sprintf("%d %s", t.degree, t.unit) }

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

var sus = map[SpeedUnit]string{KmPerHour: "km/h", MilesPerHour: "mph"}

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string { return sus[su] }

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string { return fmt.Sprintf("%d %s", s.magnitude, s.unit) }

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
// San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
func (md MeteorologyData) String() string {
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		md.location,
		md.temperature,
		md.windDirection,
		md.windSpeed,
		md.humidity,
	)
}
