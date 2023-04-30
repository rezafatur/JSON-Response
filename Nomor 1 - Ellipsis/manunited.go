package main

import (
    "fmt"
)

func average(nums ...int) float64 {
    total := 0
	
    for _, num := range nums {
        total += num
    }
    
    return float64(total) / float64(len(nums))
}

func main() {
    points := []int{3, 3, 3, 0, 1}
    
    fmt.Println("Manchester United Last 5 Matches (Premier League)")
    fmt.Println("Average Point : ", average(points...))
}
