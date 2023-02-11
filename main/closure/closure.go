package main

import "fmt"

func main() {
	/*闭包：匿名函数 + 上下文*/
	name := "hzb"
	f := func() {
		fmt.Println("My name is %s", name)
	}
	f()

	fn := ReturnClosure(name)
	fmt.Println(fn())

}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			/*延迟绑定*/
			fmt.Printf("hello, this is : %d \n", i)
		})
	}
	for _, fn := range fns {
		fn()
	}
}
