package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
	handlers map[string]func(c *Context)
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handler) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
