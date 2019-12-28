package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

var (
	xiaoli = Stu{
		Name: "小丽",
	}
	xiaoming = Stu{
		Age: 1,
	}
	xiapp = &Stu{
		Name: "p",
		Age:  1,
	}
)

func main() {
	xiaoMing := Stu{
		Name: "小明",
		Age:  12,
	}

	p := &xiaoMing
	p.Age = 10

	fmt.Println(p.Name)
	fmt.Println(p.Age) //结构体字段可通过结构体指针访问
	fmt.Println(xiaoMing.Age)

	fmt.Println(xiaoli, xiaoming, xiapp,xiaoMing)
}
