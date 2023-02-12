package main

import (
	"fmt"
	"sync"
)

func main() {

	/*指定容量创建映射*/
	m := make(map[string]string, 2)
	fmt.Printf("m: %v\n", m)

	/*不指定容量创建映射*/
	m1 := make(map[string]string)
	fmt.Printf("m1: %v\n", m1)

	/*直接初始化*/
	m2 := map[string]string{"name": "hzb"}

	_, ok := m2["name"]
	if !ok {
		println("key not found")
	}

	/*遍历顺序不确定*/
	for key, value := range m2 {
		fmt.Printf("key: %s value %s", key, value)
	}

	/*线程安全*/
	syncMap := sync.Map{}
	syncMap.Store("name", "hzb")
	value, ok := syncMap.Load("name")
	if ok {
		/*类型断言，如果值为空将会断言失败，区别于类型转换，编译器不会帮忙检查是否能断言*/
		fmt.Println(len(value.(string)))
	}
}
