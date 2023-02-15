package main

type handlerFunc func(c *Context)

type Routable interface {
	Route(method string, path string, handlerFunc handlerFunc) error
}

type Handler interface {
	Routable
	ServeHTTP(c *Context)
}
