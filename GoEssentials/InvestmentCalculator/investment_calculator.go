package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	// fmt.Print("Enter the investment amount: ")
	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Enter the expected return rate: ")
	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter the number of investment years: ")
	outputText("Enter the number of investment years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow((1+inflationRate/100), years)
	return fv, rfv
	// return
}
