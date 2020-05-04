package main

import (
	"fmt"
	"strings"
)

//闭包
func main() {
	f1 := adder1()
	fmt.Println(f1(2)) //12
	fmt.Println(f1(2)) //14
	fmt.Println(f1(2)) //16

	f2 := makeSuffixFunc(".jpg")
	f3 := makeSuffixFunc(".txt")

	fmt.Println(f2("hello"))
	fmt.Println(f3("hello"))

	f4, f5 := calc6(10)

	fmt.Println(f4(1), f5(2)) //11 9   +1 -2
	fmt.Println(f4(3), f5(4)) //12 8   +3 -4
}

//闭包
func adder1() func(int) int {
	x := 10
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc6(x int) (sum, sub func(int) int) {
	sum = func(i int) int {
		x += i
		return x
	}

	sub = func(i int) int {
		x -= i
		return x
	}

	return
}
