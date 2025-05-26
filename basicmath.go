package main

import m "math"

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
