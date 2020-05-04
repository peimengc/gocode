package main

import "fmt"

//匿名函数

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(10, 20))

	//匿名函数后面直接(参数) 直接调用
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
