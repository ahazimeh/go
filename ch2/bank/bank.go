package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
  data, _ := os.ReadFile(accountBalanceFile)
  balanceText := string(data)
  balance, _ := strconv.ParseFloat(balanceText, 64)
  return balance
}

func writeBalanceToFile(balance float64) {
  balanceText := fmt.Sprint(balance)
  os.WriteFile(accountBalanceFile, []byte(balanceText),0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

  fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposite: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositeAmount
			fmt.Println("Ballance updated! New amount:", accountBalance)
      writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't Withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Ballance updated! New amount:", accountBalance)
      writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
      break;
		}
	}

  fmt.Println("Thanks for choosing our bank")
}
