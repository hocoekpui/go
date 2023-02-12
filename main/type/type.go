package main

import (
	"net/http"
)

/*接口*/
type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)

	Start(address string) error
}

/*结构*/
type sdkHttpServer struct {
	Name string
}

/*全新类型定义，不能调用原类型的方法，需要类型转换*/
type Header map[string][]string

/*类型别名，能调用原类型的方法*/
type AliasHeader = map[string][]string

/*函数类型*/
type myFunc func(int) int
