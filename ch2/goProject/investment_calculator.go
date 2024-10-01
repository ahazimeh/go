package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

  fmt.Print("investment Amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Print("Expected Return Rate: ")
  fmt.Scan(&expectedReturnRate)

  fmt.Print("Years: ")
  fmt.Scan(&years)

	futrueValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
  futureRealValue := futrueValue / math.Pow(1+inflationRate/100,years)

  formattedFV := fmt.Sprintf("Future Value: %.4f\n", futrueValue)
  formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %f\n", futureRealValue)
  /* fmt.Println("Future Value:", futrueValue) */
  /* fmt.Printf("Future Value: %.4f\nFuture Value (adjusted for Inflation): %f\n", futrueValue, futureRealValue) */
	/* fmt.Println("Future Value (adjusted for Inflation):", futureRealValue) */
  fmt.Print(formattedFV, formattedRFV)
}

func Add(a int, b int) {
}
