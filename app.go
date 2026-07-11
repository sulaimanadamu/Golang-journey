package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 21
	result := multiply(a, b)
	div := divide(a, b)
	sub := sub(a, b)
	fmt.Printf("The multiplication of %d and %d is: %d \n", a, b, result)
	fmt.Printf("The division of %d and %d is: %d \n", a, b, div)
	fmt.Printf("The subtraction of %d from %d is: %d \n", a, b, sub)
}
