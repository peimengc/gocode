package main

import "fmt"

//函数类型与变量

//定义函数类型
type a func(int, int) int

func main() {
	var aFunc a

	fmt.Printf("%T\n", aFunc)

	aFunc = add

	fmt.Printf("%T\n", aFunc)

	re1 := aFunc(1, 2)

	fmt.Println(re1)

	aFunc1 := add

	fmt.Printf("%T\n", aFunc1)

	fmt.Println(aFunc1(1, 4))
}

func add(x, y int) int {
	return x + y
}
