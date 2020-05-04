package main

import "fmt"

func main() {
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Println("A")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B", err)
		}
	}()
	panic("panic B")
}

func funcC() {
	fmt.Println("C")
}
