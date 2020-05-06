package main

//构造函数
import "fmt"

type person1 struct {
	name string
	age  int
}

//结构体的匿名字段
type person2 struct {
	string
	int
}

func main() {
	p1 := newPerson("张三", 19)
	p1.say("你好!!")
	fmt.Println(p1)
	p1.say1("你好!!")
	fmt.Println(p1)

	p2 := &person2{
		string: "111",
		int:    0,
	}

	fmt.Println(p2.string)
}

func newPerson(name string, age int) *person1 {
	return &person1{
		name: name,
		age:  age,
	}
}

//方法和接收者
/**
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
*/
//值类型的接收者
func (p person1) say(say string) {
	p.age = 10 //无法修改原值
	fmt.Println(say, p.name)
}

//指针类型的接收者
func (p *person1) say1(say string) {
	p.age = 10
	fmt.Println(say, p.name)
}
