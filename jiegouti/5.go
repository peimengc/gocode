package main

import "fmt"

type person1 struct {
	name string
	age  int
}

func main() {
	p1 := newPerson("张三", 19)
	fmt.Println(p1)
}

func newPerson(name string, age int) *person1 {
	return &person1{
		name: name,
		age:  age,
	}
}
