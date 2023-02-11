package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))

	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *Handler
}

func (s *sdkHttpServer) Route(method, pattern string, handlerFunc func(ctx *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{Name: name, handler: &Handler{handlers: make(map[string]func(c *Context))}}
}

func index(ctx *Context) {
	fmt.Fprintf(ctx.W, "index!")
}

func main() {
	server := NewHttpServer("server")
	server.Route("GET", "/index", index)
	server.Start(":8080")
}
