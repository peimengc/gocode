package main

import "fmt"

//嵌套结构体
type Address struct {
	Code int
	City string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

type User1 struct {
	Name    string
	Age     int
	Address //嵌套匿名结构体
}

//嵌套结构体的字段名冲突
type User2 struct {
	Name    string
	Age     int
	Code    int
	Address //嵌套匿名结构体
}

func main() {

	address := &Address{
		Code: 123456,
		City: "运城",
	}

	user := &User{
		Name:    "user1",
		Age:     19,
		Address: *address,
	}

	fmt.Printf("%#v\n", user)

	user1 := &User1{
		Name:    "user2",
		Age:     29,
		Address: *address,
	}

	//访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
	fmt.Println(user1.City)

	user3 := &User2{
		Name:    "user3",
		Age:     19,
		Code:    1,
		Address: *address,
	}

	fmt.Println(user3.Code)
	fmt.Printf("%v\n", user3.Address.Code)
}
