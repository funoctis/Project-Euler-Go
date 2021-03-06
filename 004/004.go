package main

import (
	"fmt"
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
		s := 100
		e := 999

		for i := e; i > s; i-- {
			for j := e; j > i; j-- {
				t := i * j
				if t%11 != 0 {
					continue
				}
				if isPalindrome(t) && t < n {
					if t > r {
						r = t
					}
				}
			}
		}
		fmt.Println(r)
	}
}
