package main

import "fmt"

func main() {
	//无返回值函数
	sayHello()
	re := intSum(1, 3)
	fmt.Println(re)

	re2 := intSum2(1, 3)
	fmt.Println(re2)

	re3 := intSum3(10, 20, 30)
	fmt.Println(re3)

	re4 := intSum4(100, 20, 30)
	fmt.Println(re4)

	re5, re6 := intSum5(100, 20)
	fmt.Println(re5, re6)

	re7, re8 := intSum6(100, 20)
	fmt.Println(re7, re8)
}

//无返回值函数
func sayHello() {
	fmt.Println("Say Hello")
}

//固定参数函数
func intSum(x int, y int) int {
	return x + y
}

//固定参数,类型相同参数函数
func intSum2(x, y int) int {
	return x + y
}

//可变参数函数
func intSum3(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

//可变参数,固定参数混合
func intSum4(x int, y ...int) int {
	sum := x

	for _, v := range y {
		sum += v
	}

	return sum
}

//多返回值函数
func intSum5(x, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

//返回值命名
func intSum6(x, y int) (sum, sub int) {
	sub = x - y
	sum = x + y
	return
}
