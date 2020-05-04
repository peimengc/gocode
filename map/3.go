package main

import "fmt"

func main() {
	//元素为map类型的切片
	m1 := make([]map[string]string, 3, 100)

	for k, v := range m1 {
		fmt.Printf("%d,%v\n", k, v)
	}

	m1[0] = make(map[string]string, 10)
	m1[0]["name"] = "张三"
	m1[0]["age"] = "19"

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
