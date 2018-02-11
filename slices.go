package main

import "fmt"
//slice包含第一个元素但是会排除最后一个
func main() {
	primes := [6]int{2, 4, 6, 8, 10, 12}
	var s []int = primes[1:4]
	fmt.Println(s)
}
