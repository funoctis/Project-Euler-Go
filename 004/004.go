package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		var r int
		s := int(math.Pow10(n - 1))
		e := int(math.Pow10(n)) - 1

		for i := e; i > s; i-- {
			for j := e; j > i; j-- {
				t := i * j
				if (r < t) && isPalindrome(t) {
					r = t
					break
				} else if t < r {
					break
				}
			}
		}
		fmt.Println(r)
	}
}
