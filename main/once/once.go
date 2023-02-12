package main

import (
	"fmt"
	"sync"
)

/*只会执行一次，一般用于做初始化工作*/
var once sync.Once

func PrintOnce() {
	/*注意如果变量声明为局部变量将会失效*/
	once.Do(func() {
		fmt.Println("Print once")
	})
}

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}
