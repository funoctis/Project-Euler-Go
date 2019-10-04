package main

import (
    "fmt"
)

func SQsum (n int) int64 {
    var sum int64 = 0
    for i := 1; i <= n; i++ {
        sum += int64(i)
    }
    return (sum * sum)
}
func SUMsq (n int) int64 {
    var sum int64 = 0
    for i := 1; i <= n; i++ {
        sum += int64(i * i)
    }
    return sum
} 

func main() {
    var t int // Number of test cases
    fmt.Scanf("%d", &t)
    if t >= 1 && t <= 10000 {
        for i := 0; i < t; i++ {
            var n int // Each test case
            fmt.Scanf("%d", &n)
            if n >= 1 && n <= 10000 {
                sqsum := SQsum(n)
                sumsq := SUMsq(n)
                fmt.Println(sqsum - sumsq)
            }
        }
    }
}

