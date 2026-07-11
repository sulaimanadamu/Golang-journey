package main

import (
	"first-app/multiply"
	"fmt"

	//"first-app/sum"
	"first-app/division"
	"first-app/subtraction"
)

func main() {
	a := 100
	b := 21
	result := multiply.Multiply(a, b)
	div := divide.Divide(a, b)
	sub := subtract.Sub(a, b)
	fmt.Printf("The multiplication of %d and %d is: %d \n", a, b, result)
	fmt.Printf("The division of %d and %d is: %d \n", a, b, div)
	fmt.Printf("The subtraction of %d from %d is: %d \n", a, b, sub)
}
