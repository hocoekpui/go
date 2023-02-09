package main

import "fmt"

func main() {
	eat("苹果")
}

func eat(name string) {
	switch name {
	case "苹果":
		fmt.Println("不好吃")
	case "榴莲", "大蒜":
		fmt.Println("真香")
	default:
		fmt.Println("能吃么")
	}
}
