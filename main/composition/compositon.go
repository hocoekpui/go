package main

import "fmt"

type Parent struct {
}

func (p Parent) hello() {
	fmt.Println("I am " + p.getName())
}

func (p Parent) getName() string {
	return "Parent"
}

type Son struct {
	Parent
}

func (s Son) getName() string {
	return "Son"
}

type ParentInterface interface {
	do()
}

type SonInterface interface {
	ParentInterface
}

func main() {
	son := Son{Parent{}}
	son.hello()
	fmt.Println(son.getName())
	fmt.Println(son.Parent.getName())
}
