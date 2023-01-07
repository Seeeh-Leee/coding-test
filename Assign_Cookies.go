// source link: https://leetcode.com/problems/assign-cookies/
package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1,2,9}
	s := []int{1,2,9}

	sum := findContentChildren(g, s)
	fmt.Println(sum)
}

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    i := 0
    j := 0
    sum := 0

    for i < len(g) && j < len(s) {
        if s[j] >= g[i]{
            i++
            sum ++
        }
        j ++
    }

	return sum
}
