package main

import "fmt"

type Stus struct {
	Name string
	Age  int
}

var m map[string]Stus

func main() {
	m = make(map[string]Stus)

	m["class1"] = Stus{
		"张三", 0,
	}

	fmt.Println(m["class1"])

	m1 := make(map[string]int)

	m1["Answer"] = 42
	fmt.Println("The value。", m1["Answer"])

	m1["Answer"] = 81
	fmt.Println("The value.", m1["Answer"])

	delete(m1,"Answer")
	fmt.Println("The value.", m1["Answer"])

	v,ok := m1["Answer"]
	fmt.Println(v,ok)
}
