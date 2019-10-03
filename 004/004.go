package main

import (
    "fmt"
    "strconv"
)

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func isPallindrome(n string) bool {
    return n == reverse(n)
}

func main() {
    var t int
    fmt.Scanf("%d", &t)
    for z := 0; z < t; z++ {
        var n int
        fmt.Scanf("%d", &n)
        var longest_pallindrome int

        for i := 100; i <= 999; i++ {
            for j := i; j <= 999; j++ {
                var k = i*j
                if k < n && k > longest_pallindrome && isPallindrome(strconv.Itoa(k)) {
                    longest_pallindrome = k
                }
            }
        }
        fmt.Println(longest_pallindrome)
    }
}
