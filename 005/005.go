package main
import "fmt"

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scanf("%d", &n)
        if n == 1 {
            fmt.Println(1)
            continue
        }
        answer := lcm(1, 2)
        for j := 2; j <= n - 1; j++ {
            answer = lcm(answer, j + 1)
        }
        fmt.Println(answer)
    }
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}