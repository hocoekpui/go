package main

import "fmt"

func main() {
	const External = "包外可访问"
	fmt.Println(External)

	const local = "包内可访问"
	fmt.Println(local)
}
