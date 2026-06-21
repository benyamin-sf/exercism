// Package weather provide functions and variables that reports the weather and the location that is passed to its Forecast function.
package weather

var (
    // CurrentCondition represents, well... the current weather condition! Surprise 🙀.
	CurrentCondition string
    // CurrentLocation is a bit hard to explain. It's the current location?! Who would've guessed Mr. Pres, huh? 🙄.
	CurrentLocation  string
)

// Forecast function shows the CurrentCondition in the CurrentLocation (for more info on those super complicated variables, check their respective comments).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
