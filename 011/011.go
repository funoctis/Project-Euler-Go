package main
import "fmt"


var g [20][20]int

func max (x, y int) int {
    if x > y {
        return x
    }
    return y
}

func comp(x, y int) int {
    var m int
    if y + 3 < 20 {
        m = max(m, g[x][y] * g[x][y + 1] * g[x][y + 2] * g[x][y + 3])
    }
    if x + 3 < 20 {
        m = max(m, g[x][y] * g[x + 1][y] * g[x + 2][y] * g[x + 3][y])
    }
    if x + 3 < 20 && y + 3 < 20 {
        m = max(m, g[x][y] * g[x + 1][y + 1] * g[x + 2][y + 2] * g[x + 3][y + 3])
    }
    if x + 3 < 20 && y - 3 >= 0 {
        m = max(m, g[x][y] * g[x + 1][y - 1] * g[x + 2][y - 2] * g[x + 3][y - 3])
    }
    return m
}

func main() {
    for i := 0; i < 20; i++ {
        for j := 0; j < 20; j++ {
            fmt.Scanf("%d", &g[i][j])
        }
    }
    var m int
    for i := 0; i < 20; i++ {
        for j := 0; j < 20; j++ {
            m = max(m, comp(i, j))
        }
    }
    fmt.Println(m)
}

