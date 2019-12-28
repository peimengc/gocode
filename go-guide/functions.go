package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a, b := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			a(i),
			b(-2*i),
		)
	}
}
