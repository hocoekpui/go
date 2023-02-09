package main

import "fmt"

func main() {
	calculation(5, 10)
}

func calculation(begin int, end int) {
	if distance := end - begin; distance > 0 {
		fmt.Println("too far")
	} else {
		fmt.Println("slap me")
	}
}
