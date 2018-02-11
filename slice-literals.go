package main

import "fmt"

func main() {
	q := []int{2, 3, 4, 7, 9, 1}
	fmt.Println(q)
	r := []bool{true, false, true, false, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{7, false},
		{9, false},
		{1, true},
	}
	fmt.Println(s)
}
