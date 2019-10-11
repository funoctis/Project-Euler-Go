package main
import "fmt"

type ll int64

func max(a, b ll) ll {
    if a > b {
        return a
    }
    return b
}

func main() {
    var t int   // Number of test cases
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        var n ll   // Each test case value
        fmt.Scanf("%d", &n)
        var ok ll = -1
        var i, j, k ll
        for i = 1; i < n; i++ {
            j = (n * n - 2 * n * i)/(2 * n - 2 * i)
            k = n - i - j
            if (i * i + j * j == k * k && j > 0 && k > 0) {
                ok = max(ok, i * j * k)
            }
        }
        fmt.Println(ok)
    }
}
