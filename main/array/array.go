package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3}
	fmt.Printf("array: %v, len: %d, cap: %d", array, len(array), cap(array))
}
