package main

import "fmt"

func main() {
	/*var name type = value*/

	/*首字母大小写控制可访问性*/
	var External string = "包外可访问"
	fmt.Println(External)

	var local string = "包内可访问"
	fmt.Println(local)

	partial := "局部可访问"
	fmt.Println(partial)
}
