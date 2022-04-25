package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	fmt.Println(planets[2])

	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	//You can ask the Go Compiler to count the number of elements in the composite literal by specifying the ellipsis ... instead of a number.  The planets array
	// in the following listing still has a fixed length

	myPlanets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(myPlanets))
	fmt.Println(myPlanets[2])

}
