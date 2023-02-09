package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\n", array[i])
	}

	for index, value := range array {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	index := 1
	for index <= 3 {
		fmt.Println(index)
		index++
	}
}
