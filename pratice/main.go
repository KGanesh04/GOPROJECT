package main

import (
	"fmt"
	"pratice/calc"
)

type employee struct {
	id    int
	Name  string
	batch string
}

func main() {

	var num1, num2 int

	num3 := 2
	num4 := 3

	fmt.Println("num1:")
	fmt.Scan(&num1)

	fmt.Println("num2:")
	fmt.Scan(&num2)

	sub := num1 - num2

	addition := num3 + num4

	fmt.Println("a+b:", addition)

	fmt.Printf("num1 %d - num2 %d: %d", num1, num2, sub)

	fmt.Println("Hello")
	result := calc.Sum(5, 3)
	fmt.Println("addition:", result)

	var x int

	fmt.Println("enter x value")

	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Printf("x is even")
	} else {
		fmt.Printf("x is odd")
	}

	for x := 0; x < 4; {
		fmt.Println("value:", x)
		x++
	}

	E1 := employee{id: 1,
		Name:  "bob",
		batch: "developer"}

	E2 := employee{id: 2, Name: "alex", batch: "devops"}

	fmt.Println("E1Name: ", E1.Name, "-", "batch: ", E1.Name, E1.batch)
	fmt.Println("E2Name: , batch: ", E2.Name, E2.batch)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Print("array list:", arr)

}
