package main

import "fmt"

func main() {
	/*切片的声明*/
	slice := []int{1, 2, 3}
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	/*一个参数，长度与容量相同*/
	slice = make([]int, 3)
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	/*两个参数，第一个参数是长度，第二个参数是容量*/
	slice = make([]int, 3, 4)
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	/*添加元素*/
	newSlice := append(slice, 1)
	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", newSlice, len(newSlice), cap(newSlice))

	/*扩容*/
	newSlice = append(newSlice, 2)
	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", newSlice, len(newSlice), cap(newSlice))

	/*为节省空间，推荐写法*/
	lastSlice := make([]int,0,3)
	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", lastSlice, len(lastSlice), cap(lastSlice))

}
