package main

import "fmt"

func main() {
	var arr [2]string

	arr[1] = "a"
	arr[0] = "b"
	fmt.Println(arr[0],arr[1])

	fmt.Println(arr)
}