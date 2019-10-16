package main

import "fmt"

var sumPrimes [2000001]int

func check(x int) {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return
		}
	}
	sumPrimes[x] = x
}

func main() {
	for i := 2; i <= 2000000; i++ {
		check(i)
		sumPrimes[i] += sumPrimes[i-1]
	}
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(sumPrimes[n])
	}
}
