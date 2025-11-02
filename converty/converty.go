package converty

func Celsius_to_fahrenheit(celsius_value float64) float64 {
	var fahrenheit_value float64
	fahrenheit_value = celsius_value * 9 / 5 + 32
	return fahrenheit_value
}

func Fahrenheit_to_celsius(value float64) float64 {
	// var celsius_value float64
	// celsius_value = (fahrenheit_value - 32) * 5 / 9
	// return celsius_value
	return ((value - 32) * 5 / 9)
}
