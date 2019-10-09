package main
import "fmt"

var n, k int	// Number of digits and length of substring

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func mul(s string, start int) int {
    var ret int = 1
    for i := start; i < start + k; i++ {
        ret *= int(s[i] - '0')
    }
    return ret
}

func main() {
    var t int   // Number of test cases
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        fmt.Scanf("%d %d", &n, &k)
        var s string	// Number accepted as a string
        fmt.Scanf("%s", &s)
        var res int = 0
        for i := 0; i < n - k; i++ {
            res = max(res, mul(s, i))
        }
        fmt.Println(res)
    }
}