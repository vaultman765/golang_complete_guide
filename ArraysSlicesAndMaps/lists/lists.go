package lists

import "fmt"

type product struct {
	title string
	id    int
	price float64
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	h1, h2, h3 := getUserHobbies()
	hobbies := createHobbyArray(h1, h2, h3)
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	fmt.Println(hobbies[0])

	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice1 = slice1[1:3]
	fmt.Println(slice1)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Learn Go", "Build a project"}
	fmt.Println(goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Build a web app"
	fmt.Println(goals)
	goals = append(goals, "Learn more about Go")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	products := []product{
		{"Product 1", 1, 9.99},
		{"Product 2", 2, 19.99},
	}
	fmt.Println(products)

	//		Then add a third product to the existing list of products.
	products = append(products, product{"Product 3", 3, 29.99})
	fmt.Println(products)
}

func createHobbyArray(hobby1, hobby2, hobby3 string) [3]string {
	// Create a new array that contains three hobbies you have
	hobbies := [3]string{hobby1, hobby2, hobby3}
	return hobbies
}

func getUserHobbies() (string, string, string) {
	// Get user input for hobbies
	fmt.Print("Enter your first hobby: ")
	var hobby1 string
	fmt.Scanln(&hobby1)
	fmt.Print("Enter your second hobby: ")
	var hobby2 string
	fmt.Scanln(&hobby2)
	fmt.Print("Enter your third hobby: ")
	var hobby3 string
	fmt.Scanln(&hobby3)
	return hobby1, hobby2, hobby3
}
