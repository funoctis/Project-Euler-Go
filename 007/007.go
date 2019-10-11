package main
import (
    "fmt"
)

func main() {
    var ar [1e6]int
    ar[0] = 1
    ar[1] = 1
    for i := 2; i < 1e3; i++ {
        if ar[i] == 1 {
            continue
        }
        for j := i * i; j < 1e6; j += i {
            ar[j] = 1
        }
    }

    var t int   // number of test cases
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        var n int   // each test case
        fmt.Scanf("%d", &n)
        var count int = 0
        for i := 2; ; i++ {
            if ar[i] == 0 {
                count++
            }
            if count == n {
                fmt.Println(i)
                break;
            }
        }
    }
}

