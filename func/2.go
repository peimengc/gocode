package main

import "fmt"

//全局变量
var sum int

func main() {
	test()
	fmt.Println(sum)

	test1()
	//fmt.Println(x) //x 局部变量

	test2(11)
}

func test() {
	sum = 20
}

func test1() {
	x := 100
	fmt.Println(x)
}

//代码块局部变量
func test2(x int) {
	if x > 10 {
		z := 10
		fmt.Println(z)
	}

	//fmt.Println(z) //局部变量
}
