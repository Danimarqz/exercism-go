// Package weather is used to forecast weather in a given city and condition.
package weather

var (
	// CurrentCondition is used to tell the user the current weather.
	CurrentCondition string
	// CurrentLocation is usted to tell the user the current city.
	CurrentLocation  string
)

// Forecast returns a string saying the current condition of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
