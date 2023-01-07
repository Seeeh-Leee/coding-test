// source link: https://leetcode.com/problems/plus-one/
package main

import (
	"fmt"
	"math"
)

func main() {
	digits := []int{1,2,9}
	
	result := plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
