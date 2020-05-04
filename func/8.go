package main

import "fmt"

func main() {
	x := 1
	y := 2
	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	defer calc("AA", x, calc("A", x, y)) // calc("AA",1,3)
	x = 10
	defer calc("BB", x, calc("B", x, y)) // calc("BB",10,12)
	y = 20
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
