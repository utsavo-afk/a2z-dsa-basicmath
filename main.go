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
		fmt.Println("  6 → Find GCD (Greatest Common Divisor)")
		fmt.Println("  7 → Check Armstrong Number")
		fmt.Println("  8 → Find Divisors")
		fmt.Println("  9 → Find Divisors (optimized)")
		fmt.Println(" 10 → Check Prime Number")
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
		case 6:
			fmt.Println("\n🧮 Find GCD (Greatest Common Divisor)")
			fmt.Print("Enter first number: ")
			scanner.Scan()
			num1, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Print("Enter second number: ")
			scanner.Scan()
			num2, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			gcd := FindGCD(num1, num2)
			fmt.Printf("The GCD of %d and %d is: %d\n", num1, num2, gcd)
		case 7:
			fmt.Println("\n🔢 Check Armstrong Number")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			if CheckArmstrongNumber(num) {
				fmt.Printf("✅ %d is an Armstrong number.\n", num)
			} else {
				fmt.Printf("❌ %d is not an Armstrong number.\n", num)
			}
		case 8:
			fmt.Println("\n🔢 Find Divisors")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Printf("Divisors of %d: ", num)
			divisors := FindDivisors(num)
			if len(divisors) == 0 {
				fmt.Println("No divisors found.")
			} else {
				for _, divisor := range divisors {
					fmt.Print(divisor, " ")
				}
				fmt.Println()
			}
		case 9:
			fmt.Println("\n🔢 Find Divisors (optimized)")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			fmt.Printf("Divisors of %d: ", num)
			divisors := FindDivisors2(num)
			if len(divisors) == 0 {
				fmt.Println("No divisors found.")
			} else {
				for _, divisor := range divisors {
					fmt.Print(divisor, " ")
				}
				fmt.Println()
			}
		case 10:
			fmt.Println("\n🔍 Check Prime Number")
			fmt.Print("Enter a number: ")
			scanner.Scan()
			num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("❌ Invalid number.")
				continue
			}
			if CheckPrime(num) {
				fmt.Printf("✅ %d is a prime number.\n", num)
			} else {
				fmt.Printf("❌ %d is not a prime number.\n", num)
			}
		case 0:
			fmt.Println("👋 Exiting BasicMathFuncs CLI. Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please select a valid option.")
		}
	}
}
