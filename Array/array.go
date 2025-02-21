package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", slice)

	slice = append(slice, 6, 7)
	fmt.Println("Updated Slice:", slice)

	p1 := Person{Name: "Alice",
		Age:   25,
		Email: "alice@example.com"}

	p2 := Person{Name: "Bob",
		Age:   30,
		Email: "bob@example.com"}

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 1 Name:", p1.Name)

	people := []Person{p1, p2}
	for _, person := range people {
		fmt.Println("Person:", person.Name, "-", person.Age, "years old")
	}

}
