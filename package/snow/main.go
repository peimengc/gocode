package main

import (
	"fmt"
	"study/package/calc"
)

func main() {
	fmt.Println(calc.Add(10, 20))
	fmt.Println(calc.Sub(10, 20))
	fmt.Println(calc.Mul(10, 20))
	fmt.Println(calc.Into(30, 29))

	var a string = 10
	fmt.Printf(a)
}
