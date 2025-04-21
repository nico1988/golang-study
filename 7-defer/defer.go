package main

import "fmt"

func returnFunc() int {
	deferFunc()
	fmt.Println("return end")
	return 1
}

func deferFunc() {
	fmt.Println("defer end1")
	fmt.Println("defer end2")

}

func init() {
	//defer关键字
	fmt.Printf("init::hello go 1\n")
}

func main() {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	returnFunc()
}
