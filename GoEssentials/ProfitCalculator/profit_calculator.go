package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue, expenses, taxRate float64

	revenue, err1 := getUserInput("Enter the revenue: ")
	expenses, err2 := getUserInput("Enter the expenses: ")
	taxRate, err3 := getUserInput("Enter the tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1) // error messages are all the same (err1=err2=err3). Else could have specific error messages for each
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", earningsAfterTax)
	fmt.Printf("Profit ratio: %.2f\n", ratio)
	storeResults(earningsBeforeTax, earningsAfterTax, ratio)
}

func getUserInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return input, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratio = ebt / eat
	return ebt, eat, ratio
}

func storeResults(ebt, eat, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, eat, ratio)
	os.WriteFile("financials.txt", []byte(results), 0644)
}
