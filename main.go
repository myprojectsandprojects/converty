package main

import (
	"fmt"
	"github.com/myprojectsandprojects/converty/converty"
	"os"
	"strconv"
	// "log"
)

func main() {
	args := os.Args[1:]

	// fmt.Println(args)
	// fmt.Println(len(args))

	if (len(args) == 3) {
		value, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			print_usage_and_exit()
		}

		if args[1] == "Celsius" {
			if args[2] == "Fahrenheit" {
				fmt.Println(converty.Celsius_to_fahrenheit(value))
			} else {
				print_usage_and_exit()			

			}
		} else if args[1] == "Fahrenheit" {
			if args[2] == "Celsius" {
				fmt.Println(converty.Fahrenheit_to_celsius(value))
			} else {
				print_usage_and_exit()
			}
		} else {
			print_usage_and_exit()
		}
	} else {
		print_usage_and_exit()
	}
	
	// fmt.Printf("Enter Celsius value:\n")
	// var celsius_value float64
	// fmt.Scan(&celsius_value)
	// // fmt.Printf("celsius_value is %v\n", celsius_value)
	// fmt.Printf("fahrenheit value is %v\n", converty.Convert_to_fahrenheit(celsius_value))
}

func print_usage_and_exit() {
	fmt.Println("Usage:\n")
	fmt.Println("[program name] [value to be converted] [from what units] [to what units]")
	fmt.Println("Only supported units are Celsius and Fahrenheit.")
	fmt.Println("Value to be converted is a floating point number")

	os.Exit(1)
}
