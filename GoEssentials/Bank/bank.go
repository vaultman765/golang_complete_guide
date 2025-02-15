package main

import (
	"Bank/utils"
	"fmt"
)

const balanceFile = "balance.txt"

func main() {
	var accountBalance, err = utils.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter the amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount! Amount must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! Your account balance is:", accountBalance)
			utils.WriteFloatToFile(accountBalance, balanceFile)
		} else if choice == 3 {
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount! Amount must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds! Your account balance is:", accountBalance)
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Balance Updated! Your account balance is:", accountBalance)
				utils.WriteFloatToFile(accountBalance, balanceFile)
			}

		} else {
			break
		}
	}
	fmt.Println("Thank you for using Go Bank!")
}
