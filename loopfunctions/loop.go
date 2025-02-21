package main

import (
	"fmt"
)

// Function to demonstrate if, if-else, while (for loop), switch
func main() {

	// if, if-else Example
	number := 10
	if number%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// Using a while-like loop (Go does not have "while", so we use "for")
	counter := 5
	fmt.Println("Counting down:")
	for counter > 0 { // This acts as a while loop
		fmt.Println(counter)
		counter--
	}

	// switch case Example
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

}
