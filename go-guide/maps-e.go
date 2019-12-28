package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//创建map
	m := make(map[string]int)
	//获取 词 数组
	sa := strings.Fields(s)

	//循环每个词，计算每个词数量
	for _, v := range sa {
		count := m[v]
		m[v] += count + 1

		//简写
		//m[v] += 1 //因为不存在初始值为0
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
