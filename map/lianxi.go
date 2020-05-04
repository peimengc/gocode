package main

import (
	"fmt"
	"strings"
)

func main() {
	//写一个程序，统计一个字符串中每个单词出现的次数
	s := "how do you do hello"

	sSplit := strings.Split(s, " ")

	m := make(map[string]int, 100)

	for _, v := range sSplit {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	fmt.Println(m)

}
