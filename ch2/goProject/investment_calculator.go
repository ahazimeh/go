package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futrueValue = investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println(futrueValue)
}

func Add(a int, b int) {
}
