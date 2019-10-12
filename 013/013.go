package main
import (
    "fmt"
    "math/big"   
)

func main() {
    var n int   // Number of test cases
    fmt.Scanf("%d", &n)
    var num string
    bi := big.NewInt(0)
    sum := big.NewInt(0)
    for i := 0; i < n; i++ {
        fmt.Scanf("%s", &num)
        ok, _ := bi.SetString(num, 10)
        sum = sum.Add(sum, ok)   
    }
    bigString := sum.String()
    fmt.Println(bigString[:10])
}

