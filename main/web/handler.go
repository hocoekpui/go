package main

import (
	"fmt"
	"net/http"
)

type HandlerBasedOnMap struct {
	handlers map[string]func(c *Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

type Handler interface {
	http.Handler
	Routable
}

func (h *HandlerBasedOnMap) Route(method, pattern string, handlerFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

/*小技巧，可确保实现对应接口*/
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{handlers: make(map[string]func(c *Context))}
}
