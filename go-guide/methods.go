package main

import (
	"fmt"
	"math"
)

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(MyFloat(1))
	fmt.Println(f.Abs())
	fmt.Println(f.Abs1())
	fmt.Println(f)

}

type MyFloat float64

func (f *MyFloat) Abs1() float64 {
	*f += 1
	return float64(*f + 1)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
