package main

//结构体字段的可见性
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
import "fmt"

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Println(a.name, "会动")
}

type Dog struct {
	name string
	*Animal
}

func (d Dog) move() {
	if d.name != "" {
		fmt.Println(d.name, "会动1")
	} else {
		fmt.Println(d.Animal.name, "会动1")
	}
}

func main() {
	dog := Dog{
		Animal: &Animal{
			name: "大黄",
		},
	}

	dog.move()
}
