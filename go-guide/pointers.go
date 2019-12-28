package main

import "fmt"

func main() {
	i, j := 23, 123

	p := &i

	fmt.Println(*p)

	*p = 21

	fmt.Println(i)
	fmt.Println(p)

	p = &j
	fmt.Println(*p)
	*p = *p / 23

	fmt.Println(j)
}
