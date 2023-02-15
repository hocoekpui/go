package main

import "net/http"

type Server interface {
	Routable
	Start(address string)
}

type CustomerServer struct {
	filter  Filter
	handler Handler
}

func (c *CustomerServer) Route(method string, path string, handlerFunc handlerFunc) error {
	return c.handler.Route(method, path, handlerFunc)
}

func (c *CustomerServer) Start(address string) {
	err := http.ListenAndServe(address, c)
	if err != nil {
		panic(err)
	}
}

func (c *CustomerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.filter(NewContext(w, r))
}

func NewCustomerServer(filterBuilders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnTree()
	filter := handler.ServeHTTP
	for i := len(filterBuilders) - 1; i >= 0; i-- {
		filter = filterBuilders[i](filter)
	}
	return &CustomerServer{handler: handler, filter: filter}
}
