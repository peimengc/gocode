package main

import "fmt"

//指针
//任何程序数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，我们就需要指针变量。
/**
指针语法
ptr := &v    // v的类型为T

v:代表被取地址的变量，类型为T
ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
*/

func main() {
	a := 10
	b := &a
	fmt.Printf("val:%d,type:%T,&ptr:%p\n", a, &a, &a)
	fmt.Printf("val:%d,ptr:%p,ptr:%p\n", b, b, &b)
	println(&b, *&b, *b)

	x := 1
	f1(x)
	fmt.Println(x)
	//指针传值
	f2(&x)
	fmt.Println(x)
}

func f1(x int) {
	x = 100
}

func f2(x *int) {
	*x = 200
}
