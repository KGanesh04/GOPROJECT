package main

import (
	"fmt"
)

type Temparature struct {
	Value float64
}

func (t *Temparature) ToFahrenheit() float64 {
	return (t.Value * 9 / 5) + 32
}
func (t *Temparature) ToCelsius() float64 {
	return (t.Value - 32) * 5 / 9
}
func main() {
	tempF := Temparature{Value: 98.6}
	fmt.Println("Temperature in Fahrenheit:", tempF.Value)

	tempCelsius := tempF.ToCelsius()
	fmt.Println("Temperature in Celsius:", tempCelsius)

	tempC := Temparature{Value: 30}
	fmt.Println("Temperature in Fahrenheit:", tempC.Value)

	tempFahrenheit := tempF.ToFahrenheit()
	fmt.Println("Temperature in Celsius:", tempFahrenheit)
}
