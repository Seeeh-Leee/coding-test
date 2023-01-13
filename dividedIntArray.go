// source link: https://school.programmers.co.kr/learn/courses/30/lessons/12910

pakage main

import{
  "fmt"
  "sort"
}

func main {
  var arr []int{5, 9, 7, 10}
  divisor := 5
  
  result := solution(arr,divisor)
  fmt.Println(result) //[5,10]
  
}

func solution(arr []int, divisor int) []int {
    var result []int
    
    for _,value :=  range arr{
        if value % divisor==0 {
            result = append(result, value)
        }
    }
    if len(result)==0 {
        result = append(result, -1)
    }
    sort.Ints(result)
    
    return result
}

