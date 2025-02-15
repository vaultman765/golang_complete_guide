package main

import "fmt"

// Define a type alias for a map with string keys and float64 values.
type floatMap map[string]float64

// Define a method for the floatMap type.
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Allocate memory for 2 elements, but the array can grow up to 5 elements.
	// The array is initialized with 2 empty strings.
	userNames := make([]string, 2, 5)

	// Assign values to the first two elements. That are already allocated.
	// Cannot assign values to the third element, because it is not allocated yet.
	userNames[0] = "Alice"
	userNames[1] = "Bob"

	// Append values to the array. The array will grow if needed.
	// The array can grow up to 5 elements.
	userNames = append(userNames, "Max")
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	// The array is full. The next append will allocate a new array with double the capacity.
	userNames = append(userNames, "Doe")

	fmt.Println(userNames)

	// Allocate memory for 3 key-value pairs.
	// Using floatMap alias for map[string]float64.
	coursesRatings := make(floatMap, 3)

	// Assign values to the map.
	coursesRatings["Go Fundamentals"] = 5
	coursesRatings["Docker Deep Dive"] = 4.5
	coursesRatings["Kubernetes for Beginners"] = 4.7

	// Adding a new key-value pair will reallocate the map.
	coursesRatings["Python Essentials"] = 4.8

	// Call the method on the floatMap type.
	coursesRatings.output()

	// Iterate over the array.
	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range coursesRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
