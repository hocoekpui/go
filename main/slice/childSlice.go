package main

import "fmt"

func main() {
	/*切片的声明*/
	slice := []int{1, 2, 3}
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	/*子切片声明*/
	childSlice := slice[0:1]
	fmt.Printf("childSlice: %v, len: %d, cap: %d\n", childSlice, len(childSlice), cap(childSlice))

	/*子切片与原本切片共享底层数组*/
	slice[0] = 4
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("childSlice: %v\n", childSlice)

	/*结果变化不再共享数据，如扩容*/
	slice = append(slice, 4)
	slice[0] = 1
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("childSlice: %v\n", childSlice)


}
