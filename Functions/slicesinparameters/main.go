package main

func main() {
	// Function call
	numbers := []int{1, 10, 15}
	sum := sumup(&numbers)
	println(sum)

	// Variatic function call
	sumVariatic := sumupVariatic(1, 10, 15)
	println(sumVariatic)

	// Variatic function call with a slice. Ellipsis used to unpack the slice
	sumVariaticSlice := sumupVariatic(numbers...)
	println(sumVariaticSlice)
}

// A function that takes a slice of integers and returns the sum of the numbers
func sumup(numbers *[]int) int {
	sum := 0

	for _, val := range *numbers {
		sum += val
	}

	return sum
}

// Variatic function version (function that take a variable number of arguments)
func sumupVariatic(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
