package main

import "fmt"

func main() {

	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("panic content: %s\n", data)
		}
		fmt.Printf("continue")
	}()

	defer func() {
		fmt.Printf("second defer\n")
	}()

	panic("Boom")
}
