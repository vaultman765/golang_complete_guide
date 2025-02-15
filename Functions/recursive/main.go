package main

func main() {
	fact := factorial(5)
	println(fact)
}

// Recursive function to calculate the factorial of a number
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
