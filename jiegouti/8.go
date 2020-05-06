package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON序列化
//结构体标签（Tag） 格式
/**
`key1:"value1" key2:"value2"`
结构体tag由一个或多个键值对组成。
键与值使用冒号分隔，值用双引号括起来。
同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
*/

type Student struct {
	ID     int `json:"stu_id"`
	gender string
	Name   string
}

type Class struct {
	ID       int
	Title    string
	Students []*Student
}

func main() {
	c1 := Class{
		ID:       1,
		Title:    "班级1",
		Students: make([]*Student, 0, 200),
	}

	for i := 1; i < 100; i++ {
		stu := &Student{
			ID:     i,
			gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}

		c1.Students = append(c1.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c1)

	if err != nil {
		fmt.Println("json Marsha1 失败")
		return
	}

	fmt.Printf("json:%s\n", data)

	data2 := `{"ID":2,"Title":"班级2","Students":[{"ID":0,"gender":"男","Name":"stu00"},{"ID":1,"gender":"男","Name":"stu01"},{"ID":2,"gender":"男","Name":"stu02"},{"ID":3,"gender":"男","Name":"stu03"},{"ID":4,"gender":"男","Name":"stu04"},{"ID":5,"gender":"男","Name":"stu05"},{"ID":6,"gender":"男","Name":"stu06"},{"ID":7,"gender":"男","Name":"stu07"},{"ID":8,"gender":"男","Name":"stu08"},{"ID":9,"gender":"男","Name":"stu09"}]}`

	c2 := &Class{}

	err = json.Unmarshal([]byte(data2), c2)

	if err != nil {
		fmt.Println("json Unmarshal 失败")
		return
	}
}
