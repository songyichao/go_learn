package main

import (
	"fmt"
)

func fibonacci() func() int {
	f0, f1, f2, i := 0, 1, 0, 0
	return func() int {
		switch i {
		case 0:
			i++
			f2 = f0
			break
		case 1:
			i++
			f2 = f1
			break
		default:
			f2 = f1 + f0
			f0 = f1
			f1 = f2
			break
		}
		return f2
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
