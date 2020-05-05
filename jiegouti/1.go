package main

import "fmt"

func main() {
	//自定义类型
	type NewInt int

	//类型别名
	type MyInt = int

	//类型定义和类型别名的区别
	var (
		a NewInt
		b MyInt
	)

	fmt.Printf("%T\n", a) //main.NewInt
	fmt.Printf("%T\n", b) //int
	//结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。

}
