package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	sum := big.NewInt(0)
	for i := 10; i < n; i++ {
		stri := strconv.Itoa(i)
		var factDigits []*big.Int
		for _, v := range stri {
			d, _ := strconv.Atoi(string(v))
			bd := new(big.Int)
			factDigits = append(factDigits, bd.MulRange(1, int64(d)))
		}

		digSum := big.NewInt(0)
		for _, v := range factDigits {
			digSum = digSum.Add(digSum, v)
		}

		val := big.NewInt(int64(i))
		if digSum.Mod(digSum, val).String() == "0" {
			sum = sum.Add(sum, val)
		}
	}
	fmt.Println(sum)
}
