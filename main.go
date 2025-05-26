package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==================== 📐 BasicMathFuncs CLI ====================")
		fmt.Println("Choose an option:")
		fmt.Println("  1 → Count Digits")
		fmt.Println("  0 → Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		input, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("⚠️  Invalid input. Please enter a number.")
			continue
		}

		switch input {
		case 1:
			fmt.Println("\n🔢 Count Digits")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Printf("Total digits: %d\n", CountDigits(num))
		case 0:
			fmt.Println("👋 Exiting BasicMathFuncs CLI. Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please select a valid option.")
		}
	}
}
