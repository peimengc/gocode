package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//此处需要转换类型，不然会递归调用，fmt.Sprint 触发 Error
	return fmt.Sprint("不能为负数：", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func main() {
	if _, err := Sqrt(-1); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(-1))
}
