package main

import (
    "fmt"
)

func main() {
    var t int
    fmt.Scanf("%d", &t)
    if t >= 1 && t <= 100000 {
        for i := 0; i < t; i++ {
            var n int64
            fmt.Scanf("%d", &n)
            if n >= 10 && n < 40000000000000000 {
                var a, b int64 = 1, 2
                var c, sum int64
                for a < n {
                    if a % 2 == 0 {
                        sum += a
                    }
                    c = a + b
                    a = b
                    b = c
                }
                fmt.Println(sum)
            }
        }
    }
}