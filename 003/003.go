package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		var max int

		if n%2 == 0 {
			max = 2
		}
		for n%2 == 0 {
			n = n / 2
		}

		for j := 3; j*j <= n; j += 2 {
			if n%j == 0 {
				max = j
			}
			for n%j == 0 {
				n = n / j
			}
		}

		if n > 2 {
			max = n
		}

		fmt.Println(max)
	}
}
