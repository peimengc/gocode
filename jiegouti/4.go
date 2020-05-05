package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		value := stu
		m[stu.name] = &value
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	/**
	大王八 => 大王八
	小王子 => 大王八
	娜扎 => 大王八
	*/

	//根本原因在于for-range会使用同一块内存去接收循环中的值。
	/**
	因为for range创建了每个元素的副本，
	而不是直接返回每个元素的引用，
	如果使用该值变量的地址作为指向每个元素的指针，
	就会导致错误，
	在迭代时，
	返回的变量是一个迭代过程中根据切片依次赋值的新变量，
	所以值的地址总是相同的，
	导致结果不如预期。
	*/
	//https://studygolang.com/articles/9701

}
