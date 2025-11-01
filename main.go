package main

import (
	"fmt"
	"github.com/myprojectsandprojects/converty/converty"
)

func main() {
	fmt.Printf("Enter Celsius value:\n")
	var celsius_value float64
	fmt.Scan(&celsius_value)
	// fmt.Printf("celsius_value is %v\n", celsius_value)
	fmt.Printf("fahrenheit value is %v\n", converty.Convert_to_fahrenheit(celsius_value))
}
