// source link: https://leetcode.com/problems/pascals-triangle/
package main

import (
	"fmt"
)

func main() {
	numRows := 10

	sum := generate(numRows)
	fmt.Println(sum)
	// "println" print out the address of the variable, not a value. 
	// we have use "fmt.Println"
}

func generate(numRows int) [][]int {
    value := [][]int{[]int{1}} 
	fmt.Println(value) // [[1]]


    value2 := [][]int{[]int{1}, []int{1}} 
	fmt.Println(value2) // [[1] [1]]


	matrix := make([][]int, numRows)
	rows := make([]int, numRows*(numRows+1))
	for i := 0; i < numRows; i++ {
		matrix[i] = rows[0 : i+1]
	}
	fmt.Println(matrix) // [[0] [0 0] [0 0 0] [0 0 0 0] [0 0 0 0 0]]


	matrix2 := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		matrix2[i] = make([]int, i+1)
		// set 1 both end of slice
		matrix2[i][0] = 1 
        matrix2[i][i] = 1 
		for j :=1;  j < len(matrix2[i])-1; j++ {
            matrix2[i][j] = matrix2[i-1][j] + matrix2[i-1][j-1] 
        }
	}
	fmt.Println(matrix2)  //[[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]

	return matrix2
}

/* 
	for _, value := range matrix2{ 
		fmt.Println(value)
	}
*/
