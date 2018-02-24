package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pow := make([][]uint8, dy)
	for i := range pow {
		d := make([]uint8, dx)
		for l := range d {
			d[l] = uint8(i*l - 1)
		}
		pow[i] = d
	}
	return pow
}
func main() {
	pic.show(Pic)
}
