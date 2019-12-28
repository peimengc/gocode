package main

import "fmt"

type Abser interface {
	Abs() float64
}

type MyFloat1 float64

func (f *MyFloat1) Abs() float64 {
	if *f < 0 {
		return float64(-*f)
	}
	return float64(*f)
}

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return v.x * v.y
}

func main() {
	var a Abser

	f := MyFloat1(-11)

	a = &f

	fmt.Println(f)
	fmt.Println(a.Abs())
}
