package main

import "fmt"

//slice省略参数后代表数组的最大值和最小值
func main() {
	s := []int{1, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
