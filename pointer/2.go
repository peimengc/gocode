package main

//引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
/**
new与make的区别

二者都是用来做内存分配的。
make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/
import "fmt"

func main() {
	var a *int
	//*a = 10 //指针类型,未分配存储空间
	fmt.Println(a)

	b := new(int)
	*b = 10
	fmt.Println(b, *b)
}
