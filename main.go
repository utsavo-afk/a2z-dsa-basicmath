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
		fmt.Println("  2 → Reverse Number(str)")
		fmt.Println("  3 → Reverse Number(int)")
		fmt.Println("  4 → Check Palindrome")
		fmt.Println("  5 → Check Palindrome 2")
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
		case 2:
			fmt.Println("\n🔄 Reverse Number (str)")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Printf("Reversed number: %s\n", ReverseNumber(num))
		case 3:
			fmt.Println("\n🔄 Reverse Number (int)")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Printf("Reversed number: %d\n", ReverseNumber2(num))
		case 4:
			fmt.Println("\n🔍 Check Palindrome")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			if CheckPalindrome(num) {
				fmt.Printf("✅ %d is a palindrome.\n", num)
			} else {
				fmt.Printf("❌ %d is not a palindrome.\n", num)
			}
		case 5:
			fmt.Println("\n🔍 Check Palindrome 2")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			if CheckPalindrome2(num) {
				fmt.Printf("✅ %d is a palindrome.\n", num)
			} else {
				fmt.Printf("❌ %d is not a palindrome.\n", num)
			}
		case 0:
			fmt.Println("👋 Exiting BasicMathFuncs CLI. Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please select a valid option.")
		}
	}
}
