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
	root    Filter
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	var root Filter = func(c *Context) {
		handler.ServeHTTP(c)
	}

	for i := len(builders) - 1; i >= 0; i-- {
		builder := builders[i]
		root = builder(root)
	}

	return &sdkHttpServer{Name: name, handler: handler, root: root}
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
