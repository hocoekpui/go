package main

import (
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)

	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

/*全新类型定义，不能调用原类型的方法，需要类型转换*/
type Header map[string][]string

/*类型别名，能调用原类型的方法*/
type AliasHeader = map[string][]string
