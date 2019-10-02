package main

import (
	"fmt"
	"math"
	"strconv"
)

func palindrome(num int) bool {
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
		var largest int
		start := int(math.Pow10(n - 1))
		end := int(math.Pow10(n))

		for i := start; i < end; i++ {
			for j := start; j < end; j++ {
				if palindrome(i*j) && i*j > largest {
					largest = i * j
				}
			}
		}
		fmt.Println(largest)
	}
}
