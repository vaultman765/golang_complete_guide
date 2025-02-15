package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Pass an anonymous function as an argument
	transformed := transformNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(transformed)

	// Assign an anonymous function to a variable. Closure example
	double := createTransformer(2)
	triple := createTransformer(3)

	// Use the function assigned to the variable
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// A function that returns a function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
