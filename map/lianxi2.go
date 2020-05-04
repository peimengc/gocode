package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3) //s []{1,2,3}
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...) // s切片 s[:1] {1}   s[2:] {3}    追加后 s = {1,3,3}  重新赋值 s={1,3}
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

}
