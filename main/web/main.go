package main

import (
	"fmt"
	"net/http"
)

func index(c *Context) {
	_, _ = fmt.Fprintf(c.W, "index")
}

func wildcard(c *Context) {
	_, _ = fmt.Fprintf(c.W, "wildcard")
}

func param(c *Context) {
	_, _ = fmt.Fprintf(c.W, c.PathParams["param"])
}

func main() {
	server := NewCustomerServer(MetricFilterBuilder)
	_ = server.Route(http.MethodGet, "/index", index)
	_ = server.Route(http.MethodGet, "/index/*", wildcard)
	_ = server.Route(http.MethodGet, "/index/param/:param", param)

	server.Start(":8080")
}
