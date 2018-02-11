package main

import "fmt"
//slice会修改底层数组 类似引用
func main() {
	names := [4]string{
		"syc",
		"sy",
		"xwz",
		"qmn",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "xxxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
