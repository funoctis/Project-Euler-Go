package main
import "fmt"

type ll int64

func leap(y ll) bool {
    if y % 400 == 0 {
        return true
    } else if y % 100 == 0 {
        return false
    } else if y % 4 == 0 {
        return true
    } else {
        return false
    }
}

var ms = []ll{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func days(y ll, m int) ll {
    if m == 2 && leap(y) {
        return 29
    } else {
        return ms[m]
    }
}

func getDays(y ll, m int) ll {
    var (
        rem ll  = y - 1
        yls ll  = rem / 400 + rem / 4 - rem / 100
        ans ll  = yls + rem
        p   int = 1
    )

    for p < m {
        ans += days(y, p)
        p++
    }
    return ans % 7
}

func get(y ll, m int, offset ll, ys ll, ms int) ll {
    var (
        ans ll
        tmp int
    )

    if offset % 7 == 6 {
        ans = 1
    } else {
        ans = 0
    }
    tmp = ms

    for ys < y || tmp < m {
        offset = (offset + days(ys, tmp)) % 7
        if (offset % 7 == 6) {
            ans++
        }
        tmp++
        if tmp == 13 {
            tmp = 1
            ys++
        }
    }
    return ans
}

func main() {
    var t int // Number of test cases
    fmt.Scan(&t)
    for z := 0; z < t; z++ {
        var (
            y1, y2 ll
            d1, d2, m1, m2 int
        )
        fmt.Scan(&y1, &m1, &d1, &y2, &m2, &d2)
        if d1 != 1 {
            m1++
            d1 = 1
            if m1 == 13 {
                y1++
                m1 = 1
            }
        }
        var before ll = (getDays(y1, m1) - getDays(1900, 1) + 7) % 7
        var b ll = get(y2, m2, before, y1, m1)
        
        fmt.Println(b)
    }
}

