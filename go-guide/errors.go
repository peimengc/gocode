package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("在 %v,%s", e.When, e.What)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What: "出错了~~",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
