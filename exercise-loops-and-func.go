package main

import "fmt"

func Sqrt(x, lim float64) float64 {
	z := 1.0
	for ; ; {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
		y := x - z*z
		if y < 0 {
			y = -y
		}
		fmt.Println(y)
		if y <= lim {
			break
		}
	}
	return z
}
func main() {
	fmt.Println(Sqrt(5, 0.0000000001))
}
