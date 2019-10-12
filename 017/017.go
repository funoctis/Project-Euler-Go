package main
import "fmt"

type ll int64

var (
    units = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen", "Twenty"}
    decim = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
    hun string = "Hundred"
    thous string = "Thousand"
    mill string = "Million"
    bill string = "Billion"
)   

var ans []string

func solve(n ll) {
    if n <= 20 {
        ans = append(ans, units[n])
    } else if n < 100 {
        ans = append(ans, decim[n / 10])
        solve(n % 10)
    } else if n < 1000 {
        solve(n / 100)
        ans = append(ans, hun)
        solve(n % 100)
    } else if n < 1e6 {
        solve(n / 1000)
        ans = append(ans, thous)
        solve(n % 1000)
    } else if n < 1e9 {
        var h ll = 1e6
        solve(n / h)
        ans = append(ans, mill)
        solve(n % h)
    } else {
        var h ll = 1e9
        solve(n / h)
        ans = append(ans, bill)
        solve(n % h)
    }
}

func main() {
    var t int   // Number of test cases
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        var n ll
        fmt.Scanf("%d", &n)
        if n == 0 {
            fmt.Println("Zero")
        } else {
            ans = nil
            solve(n)
            fmt.Printf("%s", ans[0])
            for i := 1; i < len(ans); i++ {
                if ans[i] == "" {
                    continue
                }
                fmt.Printf(" %s", ans[i])
            }
            fmt.Printf("\n")
        }
    }
}

