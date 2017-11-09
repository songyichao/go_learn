package main

import (
	"math"
	"fmt"
)

func main() {
	const t = "Type: %T Value: %v\n"
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf(t, x, x)
	fmt.Printf(t, y, y)
	fmt.Printf(t, f, f)
	fmt.Printf(t, z, z)
}
