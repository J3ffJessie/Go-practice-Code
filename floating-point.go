package main

import "fmt"

func main() {

	third := 1.0 / 3

	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

	// The %f format verb
	// example: %4.2f  width=4 precision = .2f
	// precision specifies how many digits should appear after the decimal point; two digits for %.2f
	// width specifies the minimum number of characters to display, including the decimal point and the digits before and after the decimal.  If the width is larger than the number of characters needed, Printf will pad the left with spaces.
	// if the width is unspecified, Printf will use the number of characters necessary to display the value.const
	// To pad the left with zeros instead of spaces, prefix the width with a zero
	// fmt.Printf("%05.2f\n", third)
}
