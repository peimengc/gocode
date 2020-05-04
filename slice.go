package main

import "fmt"

func main() {
	var s1 []string
	fmt.Println(s1 == nil)

	var s2 = []int{}
	fmt.Println(s2 == nil)

	var s3 = [4]bool{true, false}
	var s4 = []bool{false, true}
	fmt.Println(s3)
	fmt.Println(s4)
	//切片是引用类型,只能跟nil比较
	//fmt.Println(s3 == s4)

	fmt.Println(len(s3), cap(s3))

	a1 := [6]int{1, 2, 3, 4, 5, 7}
	fmt.Println(a1)

	s5 := a1[0:4] //a1[low:high],（左包含，右不包含）
	fmt.Printf("s5:%v,len(s5):%v,cap(s5):%v\n", s5, len(s5), cap(s5))

	s6 := a1[:3] //省略任意索引 [:3]=[0:3] ,[3:] = [3:len(a1)]
	fmt.Println(s6)

	s51 := s5[1:6]
	fmt.Println(s51)

	//make()函数构造切片
	ms1 := make([]int, 2, 10)
	fmt.Println(ms1, len(ms1), cap(ms1))
}
