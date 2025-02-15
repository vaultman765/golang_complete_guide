package main

import (
	"fmt"
)

// Define a function type. Useful for passing functions as arguments especially with higher-order functions.
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}

	// Pass a function as an argument
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	// Pass another function as an argument
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	// Assign a function to a variable
	transformerFunction1 := getTranformerFunction(&numbers)
	transformerFunction2 := getTranformerFunction(&moreNumbers)

	// Use the function assigned to the variable
	transformedNumbers1 := transformNumbers(&numbers, transformerFunction1)
	fmt.Println(transformedNumbers1)

	transformedNumbers2 := transformNumbers(&moreNumbers, transformerFunction2)
	fmt.Println(transformedNumbers2)

}

// A function that returns a function
func getTranformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// A function that takes a slice of integers and a function as arguments
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		// Call the function passed as an argument
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
