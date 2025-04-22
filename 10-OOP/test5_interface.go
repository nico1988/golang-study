package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 改如何区分 此时引用的底层数据类型到底是什么？

	//给 interface{} 提供 “类型断言” 的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}


	switch arg.(type) {
	case int:
		fmt.Println("arg is int type")
	case string:
		fmt.Println("arg is string type")
	case float64:
		fmt.Println("arg is float64 type")
	case Book:
		fmt.Println("arg is Book type")
	default:
		fmt.Println("arg is unknown type")
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
