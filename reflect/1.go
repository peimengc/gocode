package main
/**
类型断言
x.(T)
x：表示类型为interface{}的变量
T：表示断言x可能是的类型。
 */
import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n", t.Name(), t.Kind())
	fmt.Println(t)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("kind:%v\n", v.Kind())
	fmt.Println(v,reflect.Int)
}

func main() {
	var (
		a int64
		b int32
		c float32
		d string
		e rune
		f *rune
	)
	reflectType(a)
	reflectType(b)
	reflectType(c)
	reflectType(d)
	reflectType(e)
	reflectType(f)

	g := 100
	reflectValue(g)
}
