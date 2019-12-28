package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var v float64 = 1
	for z := 1; z <= 10; z++ {
		v = v - (v*v-x)/(v*2)
		//fmt.Println(v)
	}
	return v
}

func Sqrt1(x float64) float64 {
	var a, b float64
	for a, b = 1, 0; a-b != 0; b, a = a, a-(a*a-x)/(a*2) {

	}
	return a
}

func main() {
	fmt.Println(
		math.Sqrt(5),
		Sqrt(5),
		Sqrt1(5),
	)
}
