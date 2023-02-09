package main

import "fmt"

func main() {
	name := "hzb"
	age := 24

	/*格式化字符串*/
	str := fmt.Sprintf("My name is %s, I am %d", name, age)
	fmt.Println(str)

	/*格式化字符串并输出到控制台*/
	fmt.Printf("My name is %s, I am %d", name, age)
}
