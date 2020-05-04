package main

import "fmt"

func main() {
	a := make([]string, 5, 10)

	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}

	fmt.Printf("%+v", a[0])
	fmt.Println(a, len(a), cap(a))
}
