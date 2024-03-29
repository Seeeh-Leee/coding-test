// link ; https://school.programmers.co.kr/learn/courses/30/lessons/120956?language=go

import (
    // "fmt"
    "strings"
)

func solution(babbling []string) int {
    result  := 0
    count   := 0
    var bab         = []string{"aya", "ye", "woo", "ma"}
    var babCount    = []int{3,2,3,2}
    
    for _, value := range babbling {
        count =0
        for i:=0; i<len(bab); i++{
            if (strings.Contains(value, bab[i])) {
                count += babCount[i]
            }
        }
        if (len(value) == count){
            result++
        }
    }
    return result
}
