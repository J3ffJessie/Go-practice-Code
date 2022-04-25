package main

import "fmt"

func main() {
	// Integer types, FIVE are signed, meaning they can represent both positive and negative whole numbers. the most common abbreviation is int

	//year int := 2022

	//The other FIVE are unsigned, meaning they're for positive numbers only.  The abbreviation for unsigned integer is unit

	//month uint := 2

	// When using type inference, Go will always pick the int type for a literal whole number.  The following three lines are all equivalent

	//year := 2022
	//var year = 2022
	//var year int = 2022

	//It is preferable not to specify the int type when it can be inferred

	//Inspect a variable's type:

	year := 2022
	fmt.Printf("Type %T for %v\n", year, year)

	//Instead of repeating the variable twice, you can tell Printf to use the first argument [1] for the second format verb

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

}
