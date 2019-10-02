package main

import "fmt"

func main() {
	var t int //number of test cases
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int //find upto n
		fmt.Scanf("%d", &n)
		a := (n - 1) / 3                 //count of numbers divisible by 3
		b := (n - 1) / 5                 //count of numbers divisible by 5
		c := (n - 1) / 15                //count of numbers divisible by 15
		sumOf3 := 3 * a * (a + 1) / 2    //sum of numbers divisible by 3
		sumOf5 := 5 * b * (b + 1) / 2    //sum of numbers divisible by 5
		sumOf15 := 15 * c * (c + 1) / 2  //sum of numbers divisible by 3
		sum := sumOf3 + sumOf5 - sumOf15 //answer
		fmt.Println(sum)
	}
}
