package main

import "fmt"

type Man struct {
	Name string
}

func main() {
	structMan := Man{}
	structMan.setName("hzb")
	fmt.Printf("%v\n", structMan)
	structMan.setNamePointer("hzb")
	fmt.Printf("%v\n", structMan)

	pointerMan := &Man{}
	pointerMan.setName("hzb")
	fmt.Printf("%v\n", pointerMan)
	pointerMan.setNamePointer("hzb")
	fmt.Printf("%v\n", pointerMan)
}

/*结构体接收器方法不会影响自身状态*/
func (man Man) setName(name string) {
	man.Name = name
}

/*指针接收器，影响自身状态*/
func (man *Man) setNamePointer(name string) {
	man.Name = name
}
