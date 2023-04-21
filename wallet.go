package main

import (
	"fmt" //fmt package used to print out statements
)

// main method
func main() {
	fmt.Println("Input your name : ")
	var nameInput string
	fmt.Scan(&nameInput)

	fmt.Println("Input how much you want to fund your wallet with: ")
	var fundInput int
	fmt.Scan(&fundInput)

	fmt.Println("Welcome " + nameInput + "\nInput the items you would like to purchase")

	//for loop that executes repeatedly
	for {
		fmt.Println("1. Shoes ~ $300")
		fmt.Println("2. Eggs ~ $200")
		fmt.Println("3. Chips ~ $100")
		fmt.Println("4. Dell laptop ~ $800")
		fmt.Println("5. Pringles ~ $10")
		fmt.Println("6. Check Balance")
		fmt.Println("7. Exit App")

		fmt.Println("\nEnter choice : ")
		var optionsInput int
		fmt.Scanln(&optionsInput)

		switch optionsInput {

		case 1:
			var price int = 300

			if fundInput < price {
				fmt.Println("Your wallet balance can currently not afford this item.")
			} else if fundInput <= 0 {
				fmt.Println("Your wallet is empty")
			} else {
				fundInput = fundInput - price
				fmt.Println(fundInput)
			}
		case 2:
			var price2 int = 200

			if fundInput < price2 {
				fmt.Println("Your wallet balance can currently not afford this item.")
			} else if fundInput <= 0 {
				fmt.Println("Your wallet is empty")
			} else {
				fundInput = fundInput - price2
				fmt.Println(fundInput)
			}
		case 3:
			var price3 int = 100

			if fundInput < price3 {
				fmt.Println("Your wallet balance can currently not afford this item.")
			} else if fundInput <= 0 {
				fmt.Println("Your wallet is empty")
			} else {
				fundInput = fundInput - price3
				fmt.Println(fundInput)
			}
		case 4:
			var price4 int = 800

			if fundInput < price4 {
				fmt.Println("Your wallet balance can currently not afford this item.")
			} else if fundInput <= 0 {
				fmt.Println("Your wallet is empty")
			} else {
				fundInput = fundInput - price4
				fmt.Println(fundInput)
			}

		case 5:
			var price5 int = 10

			if fundInput < price5 {
				fmt.Println("Your wallet balance can currently not afford this item.")
			} else if fundInput <= 0 {
				fmt.Println("Your wallet is empty")
			} else {
				fundInput = fundInput - price5
				fmt.Println(fundInput)
			}
		case 6:

			fmt.Println(fundInput)

		case 7:
			fmt.Println("Goodbye")
			return
			break
		default:
			fmt.Println("default case")

		}

		fmt.Println() //print statement to execute the contents of "for" repeatedly
	}
}