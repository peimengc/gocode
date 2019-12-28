package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("p==", p)
	fmt.Println("p[2:3]==", p[2:3])
	fmt.Println("p[1:3]==", p[1:3])
	fmt.Println("p[0:3]==", p[0:3])
	fmt.Println("p[1:]==", p[1:])
	fmt.Println("p[:4]==", p[:4])

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var e []int

	printSlice("e",e)

	if e == nil {
		fmt.Println("nil")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
