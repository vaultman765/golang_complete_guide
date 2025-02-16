package cmdmanager

import "fmt"

type CMDManager struct {
}

// ReadLines reads the prices from the command line and returns them as a slice of strings
func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER. Type 'done' when you are finished.")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "done" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

// WriteResult writes the result to the command line
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

// New creates a new CMDManager struct
func New() CMDManager {
	return CMDManager{}
}
