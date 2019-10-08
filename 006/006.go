package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		sum := n * (n + 1) / 2
		squareOfSum := sum * sum
		sumOfSquare := n * (2*n + 1) * (n + 1) / 6
		fmt.Println(squareOfSum - sumOfSquare)
	}
}
