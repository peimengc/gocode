package main

import (
	"fmt"
	"unsafe"
)

//结构体初始化
type person struct {
	name, sex string
	age       int
}

func main() {
	p1 := person{
		name: "p1",
		sex:  "n",
		age:  10,
	}

	p2 := &person{
		name: "p2",
		sex:  "n",
		age:  18,
	}

	p3 := person{
		name: "",
	}

	//全部, 顺序, 统一
	p4 := person{
		"p4",
		"n",
		19,
	}

	pp := &person{
		name: "",
		sex:  "",
		age:  0,
	}

	fmt.Printf("%p,%p\n", p2, pp)

	fmt.Printf("p1%#v,p2%#v,p3%#v,p4%#v\n", p1, p2, p3, p4)
	fmt.Printf("p1%v,p2%v,p3%v,p4%v\n", unsafe.Sizeof(p1), unsafe.Sizeof(p2), unsafe.Sizeof(p3), unsafe.Sizeof(p4))
}
