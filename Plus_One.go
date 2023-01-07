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
	a, count  := 0, 0
	var result []int

	if (digits[len(digits)-1] == 9){
		
		for i:= 0; i< len(digits); i++{ 
			a +=  digits[i] * 
						int(math.Pow(10, float64(len(digits)-i-1))) 
		}
		a++
		// save the result
		add_value := a 

		// count the digit of add_value
		for ; a != 0; count++ {
			a = a/10
		}

		// make dividend by the digit count
		dividend := int(math.Pow(10, float64(count)-1))  
	
		// append the value in array
		for i:=0; i<count; i++ {
			a = (add_value / dividend)%10

			dividend = dividend/10
			result = append(result,a)
		}
	} else {
		//if last digit isn't 9. just add and return
		digits[len(digits)-1]++
		result = digits
	}
	
	return result
}
