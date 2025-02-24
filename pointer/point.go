package main

import (
	"fmt"
)

// Struct with a Pointer Receiver Method
type Person struct {
	Name string
	Age  int
}

// Method that modifies struct data using a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++ // Modifies the original Person object
}

// Standalone Function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Function to demonstrate pointers
func modifyValue(x *int) {
	*x += 10 // Dereferencing the pointer to modify original value
}

// Function demonstrating a map
func getCountryCode() map[string]string {
	// Map to store country codes
	countryCodes := map[string]string{
		"US": "United States",
		"IN": "India",
		"UK": "United Kingdom",
		"CA": "Canada",
	}
	return countryCodes
}

func main() {
	//  1. Pointers
	num := 20
	ptr := &num // Pointer stores address of num
	fmt.Println("Original Value:", num)
	modifyValue(ptr) // Pass pointer to function
	fmt.Println("Modified Value:", num)

	//  2. Methods (Using Struct)
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("\nBefore Birthday:", person.Name, "is", person.Age, "years old")
	person.HaveBirthday() // Calls the method to increment age
	fmt.Println("After Birthday:", person.Name, "is", person.Age, "years old")

	//  3. Functions
	fmt.Println("\nGreeting:", greet("Bob"))

	//  4. Maps
	countryMap := getCountryCode()
	fmt.Println("\nCountry Codes:")
	for code, country := range countryMap {
		fmt.Println(code, "->", country)
	}
}
