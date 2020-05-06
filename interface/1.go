package main

import "fmt"

type Mover interface {
	move()
}

type Dog struct{}

//使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
/*func (d Dog) move() {
	fmt.Println("move")
}*/

//此时实现Mover接口的是*dog类型，所以不能给a传入dog类型的d1，此时a只能存储*dog类型的值。
func (d *Dog) move() {
	fmt.Println("move")
}

func main() {
	var a Mover

	//d1 := Dog{}
	//a = d1
	//a.move()
	d2 := &Dog{}
	a = d2
	a.move()
}
