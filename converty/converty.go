package converty

func Convert_to_fahrenheit(celsius_value float64) float64 {
	var fahrenheit_value float64
	fahrenheit_value = celsius_value * 9 / 5 + 32
	return fahrenheit_value
}

