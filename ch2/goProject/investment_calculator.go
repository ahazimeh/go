package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	/* fmt.Print("Investment Amount: ") */
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	/* fmt.Print("Expected Return Rate: ") */
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	/* fmt.Print("Years: ") */
	outputText("Years: ")
	fmt.Scan(&years)

  futrueValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.4f\n", futrueValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %f\n", futureRealValue)
	/* fmt.Println("Future Value:", futrueValue) */
	/* fmt.Printf("Future Value: %.4f\nFuture Value (adjusted for Inflation): %f\n", futrueValue, futureRealValue) */
	/* fmt.Println("Future Value (adjusted for Inflation):", futureRealValue) */
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	rfv = fv / math.Pow(1+inflationRate/100, years)
  return
	/* return fv, rfv */
}
