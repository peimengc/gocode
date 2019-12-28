package main

import "fmt"

func fibonacci() func() int {
	v1, v2 := 0, 1
	return func() int {
		temp := v1
		v1, v2 = v2, v1+v2
		return temp
	}
}

func main() {
	fi := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fi())
	}
}
