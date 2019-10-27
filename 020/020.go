package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func digitSum(n *big.Int) int {
	sum := 0
	for _, digit := range n.String() {
		d, _ := strconv.Atoi(string(digit))
		sum += d
	}
	return sum
}

func main() {
	var n int
	var values []int
	var x = new(big.Int)
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Scanf("%d", &v)
		values = append(values, v)
	}

	for _, v := range values {
		fmt.Println(digitSum(x.MulRange(1, int64(v))))
	}

}
