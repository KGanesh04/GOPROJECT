package main

import "fmt"

func main() {
	var num1, num2 int
	num3 := 4
	num4 := 5

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	sum1 := num3 + num4

	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
	fmt.Printf("sum of  num3:%d ,num4:%d is %d\n", num3, num4, sum1)
}
