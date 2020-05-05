package main

import "fmt"

/**
结构体的定义
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
*/
func main() {
	//结构体的定义
	type person struct {
		name, sex string //相同类型
		age       int
	}

	var peimeng person
	peimeng.age = 18
	peimeng.name = "裴孟"
	peimeng.sex = "男"

	fmt.Printf("%T\n", peimeng)
	fmt.Printf("%v\n", peimeng)
	fmt.Printf("%#v\n", peimeng)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}

	user.Name = "古力娜扎"
	user.Age = 18

	fmt.Printf("%#v\n", user)

	//指针结构体
	dahe := new(person) //&person{} 同等效果

	dahe.name = "大和"

	fmt.Printf("%v\n", dahe)
	fmt.Printf("%T\n", dahe)
	fmt.Printf("%#v\n", dahe)
	fmt.Printf("%p\n", dahe)

	yanling := &person{name: "言灵"}

	fmt.Printf("%#v\n", yanling)
}
