package main

import (
	m "math"
	"strconv"
)

// Loop over number and repeat division until num is 0
// func countDigits(n int) int {
// 	count := 0
// 	for n > 0 {
// 		count++
// 		n /= 10
// 	}
// 	return count
// }

func CountDigits(n int) int {
	count := int(m.Log10(float64(n)) + 1)
	return count
}

func ReverseNumber(n int) string {
	s := ""
	for n > 0 {
		d := n % 10
		s += strconv.Itoa(d)
		n /= 10
	}
	return s
}

func ReverseNumber2(n int) int {
	var revNum int
	for n > 0 {
		ld := n % 10
		revNum = (revNum * 10) + ld
		n /= 10
	}
	return revNum
}

func CheckPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
		if i == j {
			break
		}
	}
	return true
}

func CheckPalindrome2(n int) bool {
	revNum := 0
	dup := n
	for n > 0 {
		ld := n % 10
		revNum = (revNum * 10) + ld
		n /= 10
	}
	if dup == revNum {
		return true
	} else {
		return false
	}
}

func FindGCD(a, b int) int {
	gcd := 1
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}
	return gcd
}
