package main

import (
	"fmt"
	"math"
)

func main() {
	// values needed
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64= 10

	// calculation
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// result
	fmt.Println(futureValue)
}
