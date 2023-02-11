package main

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
}

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{Name: name, handler: NewHandlerBasedOnMap()}
}

func index(ctx *Context) {
	fmt.Fprintf(ctx.W, "index!")
}

func main() {
	server := NewHttpServer("server")
	server.Route(http.MethodGet, "/index", index)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
