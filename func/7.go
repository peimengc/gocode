package main

import "fmt"

//defer
/**
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
在defer归属的函数即将返回时，
将延迟处理的语句按defer定义的逆序进行执行，
也就是说，
先被defer的语句最后被执行，
最后被defer的语句，
最先被执行。
*/
func main() {
	fmt.Println("开始")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("结束")

	fmt.Printf("f1:%v\n", f1()) //5
	fmt.Printf("f2:%v\n", f2()) //6
	fmt.Printf("f3:%v\n", f3()) //5
	fmt.Printf("f4:%v\n", f4()) //5
}

func f1() int {
	x := 5
	//先赋值后defer 所以5
	defer func() { x++ }()
	return x
}

//先return后defer 所以 x=5 之后 x++ x=6
func f2() (x int) {
	//两个x是同一个
	defer func() { x++ }()
	return 5
}

//先赋值,后defer,所以y=x=5 之后x++
func f3() (y int) {
	x := 5
	defer func() { x++ }()
	return x
}

//
func f4() (x int) {
	//两个x不是同一个
	defer func(x int) { x++ }(x)
	return 5
}
