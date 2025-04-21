package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2, numbers len = 5， [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 切片扩容机制，给一个切片追加一个元素3, numbers len = 6， [0,0,0,1,2,3], cap = 10
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("-=-------")
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	// 向numbers2切片追加一个元素1, numbers2 len = 4， [0,0,0,1], cap = 3
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
