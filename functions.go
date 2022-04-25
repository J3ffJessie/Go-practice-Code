package main

import "fmt"

//kevlin to celsius converts Kelvin to Celsius

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFarenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32
	return c
}

func main() {
	kelvin := 0.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "K is ", celsius, "C")

	farenheit := celsiusToFarenheit(celsius)

	fmt.Println(celsius, "C is", farenheit, "F")
}
