package main

import (
	"fmt"
	"math"
)

func main() {
	// values needed
	var investmentAmount, years, expectedReturnRate = 1000.0, 10.0, 5.5

	// calculation
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// result
	fmt.Println(futureValue)
}
