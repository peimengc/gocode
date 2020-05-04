package main

import (
	"errors"
	"fmt"
)

//高阶函数

func main() {
	fmt.Println(calc4(10, 20, sum4))
	fmt.Println(calc4(10, 20, sub4))

	fun, err := calcFun("*")

	if err == nil {
		fmt.Println(fun(10, 20))
	} else {
		fmt.Println(err)
	}
}

//函数作为参数
func calc4(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func calcFun(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return sum4, nil
	case "-":
		return sub4, nil
	default:
		err := errors.New(fmt.Sprintf("'%s'暂不支持的计算类型", s))
		return nil, err
	}
}

func sum4(x, y int) int {
	return x + y
}

func sub4(x, y int) int {
	return x - y
}
