package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	m1 := make(map[string]int, 100)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		m1[key] = value
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//按顺序遍历map
	//1先取出所有key放入切片中
	keys := make([]string, 0, 100)

	for k := range m1 {
		keys = append(keys, k)
	}

	//key 排序
	sort.Strings(keys)

	fmt.Println(keys)

	//遍历key
	for _, v := range keys {
		fmt.Println(v, m1[v])
	}
}
