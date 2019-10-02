package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)

		a := 1
		b := 2
		var c, sum int
		for a < n {
			if a%2 == 0 {
				sum += a
			}
			c = a + b
			a = b
			b = c
		}

		fmt.Println(sum)
	}
}
