package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*字符长度*/
	fmt.Println(len("你好"))
	/*字符数量*/
	fmt.Println(utf8.RuneCountInString("你好"))
}