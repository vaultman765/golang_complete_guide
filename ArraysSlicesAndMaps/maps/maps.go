package main

import "fmt"

func main() {
	// Creating a map
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites)

	// Accessing a value
	fmt.Println(websites["Amazon Web Services"])

	// Adding a new key-value pair
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	// Deleting a key-value pair
	delete(websites, "Google")
	fmt.Println(websites)
}
