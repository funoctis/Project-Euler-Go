package main
import "fmt"

var v []int

func main() {
    var two []int
    two = append(two, 1)
    v = append(v, 1)
    var power int
    for ; power <= 1e4; {
        var digitSum int
        var carry int
        for i := 0; i < len(two); i++ {
            carry += two[i] * 2
            two[i] = carry % 10
            digitSum += two[i]
            carry /= 10
        }
        for carry != 0 {
            two = append(two, carry % 10)
            digitSum += carry % 10
            carry /= 10
        }
        v = append(v, digitSum)
        power++
    }
    var t int   // Number of test cases
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        var n int   // Each test case
        fmt.Scanf("%d", &n)
        fmt.Println(v[n])
    }
}

