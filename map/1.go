package main

import "fmt"

func main() {
	m1 := map[string]int{
		"张三": 100,
		"里斯": 99,
	}
	//v值/默认值  ok bool 是否存在
	v, ok := m1["张散"]

	fmt.Println(v, ok)
}
