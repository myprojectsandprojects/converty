package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter Celsius value:\n")
	var celsius_value float64
	fmt.Scan(&celsius_value)
	// fmt.Printf("celsius_value is %v\n", celsius_value)
	fmt.Printf("fahrenheit value is %v\n", convert_to_fahrenheit(celsius_value))
}

func convert_to_fahrenheit(celsius_value float64) float64 {
	var fahrenheit_value float64
	fahrenheit_value = celsius_value * 9 / 5 + 32
	return fahrenheit_value
}
